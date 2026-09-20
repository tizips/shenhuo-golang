// Package wechat 封装微信公众号（服务号）能力：网页授权解析 OpenID 与模板消息推送。
package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	tmplRequest "github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/herhe-com/framework/facades"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/tizips/shenhuo/model"
)

var (
	account    *officialAccount.OfficialAccount
	accountErr error
	accountOne sync.Once
)

// tokenCache 返回基于框架 Redis 的缓存实现，用于跨进程共享微信 access_token；
// Redis 未注册或连接不可用时返回 nil，由 PowerWeChat 回退到进程内默认缓存。
func tokenCache() cache.CacheInterface {

	redis, ok := facades.OptionalRedis()
	if !ok {
		return nil
	}

	client := redis.Default()
	if client == nil {
		return nil
	}

	return &redisCache{client: client}
}

// redisCache 基于 go-redis 实现 PowerLibs 的 cache.CacheInterface；
// 序列化行为与 PowerWeChat 默认的 GRedis 缓存保持一致（JSON 编解码）。
type redisCache struct {
	client *redis.Client
}

func (c *redisCache) Get(key string, defaultValue any) (any, error) {

	raw, err := c.client.Get(context.Background(), key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, cache.ErrCacheMiss
	}
	if err != nil {
		return nil, err
	}

	var value any
	if err = json.Unmarshal(raw, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func (c *redisCache) Set(key string, value any, expires time.Duration) error {

	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(context.Background(), key, raw, expires).Err()
}

func (c *redisCache) Has(key string) bool {
	value, err := c.Get(key, nil)
	return err == nil && value != nil
}

func (c *redisCache) AddNX(key string, value any, ttl time.Duration) bool {

	raw, err := json.Marshal(value)
	if err != nil {
		return false
	}

	set, err := c.client.SetNX(context.Background(), key, raw, ttl).Result()
	return err == nil && set
}

func (c *redisCache) Add(key string, value any, ttl time.Duration) error {

	if c.Has(key) {
		return errors.New("cache: key already exists")
	}

	return c.Set(key, value, ttl)
}

func (c *redisCache) Remember(key string, ttl time.Duration, callback func() (any, error)) (any, error) {

	value, err := c.Get(key, nil)
	if err == nil && value != nil {
		return value, nil
	}

	value, err = callback()
	if err != nil {
		return nil, err
	}

	if err = c.Set(key, value, ttl); err != nil {
		return nil, err
	}

	return value, nil
}

// app 获取服务号应用单例；配置缺失或初始化失败时返回错误。
func app() (*officialAccount.OfficialAccount, error) {

	accountOne.Do(func() {

		appID := facades.Config().GetString("wechat.official_account.app_id", "")
		secret := facades.Config().GetString("wechat.official_account.secret", "")
		if appID == "" || secret == "" {
			accountErr = errors.New("微信公众号未配置（wechat.official_account）")
			return
		}

		account, accountErr = officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
			AppID:  appID,
			Secret: secret,
			Cache:  tokenCache(),
			// 必须显式声明网页授权 scope：缺失时 PowerWeChat 会回退到 snsapi_login，
			// 与授权链接的 snsapi_base 不一致，导致换不到 OpenID。
			OAuth: officialAccount.OAuth{
				Scopes: []string{facades.Config().GetString("wechat.official_account.oauth.scope", "snsapi_base")},
			},
		})
	})

	return account, accountErr
}

// OAuthURL 生成微信服务号网页授权跳转地址。
// 授权后微信将重定向到配置的回调地址（wechat.official_account.oauth.redirect），
// 并附带 code 与 state 参数；scope 缺省为 snsapi_base。
func OAuthURL(state string) (string, error) {

	appID := facades.Config().GetString("wechat.official_account.app_id", "")
	if appID == "" {
		return "", errors.New("微信公众号未配置（wechat.official_account）")
	}

	redirect := facades.Config().GetString("wechat.official_account.oauth.redirect", "")
	if redirect == "" {
		return "", errors.New("微信网页授权回调地址未配置（wechat.official_account.oauth.redirect）")
	}

	scope := facades.Config().GetString("wechat.official_account.oauth.scope", "snsapi_base")

	query := url.Values{}
	query.Set("appid", appID)
	query.Set("redirect_uri", redirect)
	query.Set("response_type", "code")
	query.Set("scope", scope)
	if state != "" {
		query.Set("state", state)
	}

	return "https://open.weixin.qq.com/connect/oauth2/authorize?" + query.Encode() + "#wechat_redirect", nil
}

// OpenIDOfCode 通过微信服务号网页授权 Code 解析用户 OpenID。
// 直接使用 TokenFromCode 换取授权令牌，从响应中取 openid；
// app.debug 开启时打印完整解析内容，便于排查授权失败。
func OpenIDOfCode(c context.Context, code string) (string, error) {

	oa, err := app()
	if err != nil {
		return "", err
	}

	token, err := oa.OAuth.TokenFromCode(code)
	if err != nil {
		return "", err
	}

	if facades.Config().GetBool("app.debug", false) {
		if raw, err := json.Marshal(token); err == nil {
			fmt.Printf("[wechat] TokenFromCode: %s\n", raw)
		}
	}

	openID, _ := (*token)["openid"].(string)
	if openID == "" {
		return "", errors.New("网页授权 Code 无效，无法解析出 OpenID")
	}

	return openID, nil
}

// template 模板消息配置：模板ID、消息打开链接与关键词变量名。
type template struct {
	ID   string
	URL  string
	Keys map[string]string
}

// templateOf 读取模板配置；keys 缺省时使用微信模板实际变量名。
func templateOf(name string) (template, error) {

	prefix := "wechat.templates." + name

	pick := func(key, fallback string) string {
		if value := facades.Config().GetString(key, ""); value != "" {
			return value
		}
		return fallback
	}

	tpl := template{
		ID:  facades.Config().GetString(prefix+".id", ""),
		URL: facades.Config().GetString(prefix+".url", ""),
		Keys: map[string]string{
			"name":   pick(prefix+".keys.name", "thing2"),
			"number": pick(prefix+".keys.number", "character_string3"),
			"unit":   pick(prefix+".keys.unit", "thing3"),
		},
	}

	if tpl.ID == "" {
		return tpl, fmt.Errorf("微信模板ID未配置（%s.id）", prefix)
	}

	return tpl, nil
}

// SendScoreNotice 发送成绩发布通知（用户名称、用户编号）。
// total 为总成绩，会以 total 查询参数追加到消息打开链接中。
func SendScoreNotice(c context.Context, openID, name, number, total string) error {

	tpl, err := templateOf("score")
	if err != nil {
		return err
	}

	link := withQuery(tpl.URL, map[string]string{"name": name, "total": total})

	return send(c, openID, tpl.ID, link, &power.HashMap{
		tpl.Keys["name"]:   keyword(name),
		tpl.Keys["number"]: keyword(number),
	})
}

// SendDrawNotice 发送成绩提醒（客户名称、公司名称）。
// group 为抽签分组名称，会以 name/group 查询参数追加到消息打开链接中。
func SendDrawNotice(c context.Context, openID, name, unit, group string) error {

	tpl, err := templateOf("draw")
	if err != nil {
		return err
	}

	link := withQuery(tpl.URL, map[string]string{"name": name, "group": group})

	return send(c, openID, tpl.ID, link, &power.HashMap{
		tpl.Keys["name"]: keyword(name),
		tpl.Keys["unit"]: keyword(unit),
	})
}

// NotifyDraws 抽签结束后给中签人员发送成绩提醒。
// 通知为尽力而为：人员未绑定 OpenID 或发送失败时跳过，不影响抽签结果。
func NotifyDraws(c context.Context, db *gorm.DB, draws []model.ShDraw) {

	for _, draw := range draws {

		var person model.ShPerson
		if err := db.First(&person, "`id`=?", draw.PersonID).Error; err != nil || person.OpenID == "" {
			continue
		}

		group := ""
		if draw.Category != nil {
			group = draw.Category.Name
		}

		_ = SendDrawNotice(c, person.OpenID, person.Name, person.Unit, group)
	}
}

// withQuery 在链接上追加查询参数；参数值为空时跳过，不覆盖链接上已有的参数。
func withQuery(link string, params map[string]string) string {

	if link == "" || len(params) == 0 {
		return link
	}

	u, err := url.Parse(link)
	if err != nil {
		return link
	}

	query := u.Query()
	for key, value := range params {
		if value != "" && query.Get(key) == "" {
			query.Set(key, value)
		}
	}
	u.RawQuery = query.Encode()

	return u.String()
}

// send 发送模板消息；link 非空时点击消息将跳转到该链接。
func send(c context.Context, openID, templateID, link string, data *power.HashMap) error {

	oa, err := app()
	if err != nil {
		return err
	}

	result, err := oa.TemplateMessage.Send(c, &tmplRequest.RequestTemlateMessage{
		ToUser:     openID,
		TemplateID: templateID,
		URL:        link,
		Data:       data,
	})
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return errors.New("微信模板消息发送失败：" + result.ErrMsg)
	}

	return nil
}

// keyword 组装模板消息关键词项。
func keyword(value string) map[string]string {
	return map[string]string{"value": value, "color": "#173177"}
}
