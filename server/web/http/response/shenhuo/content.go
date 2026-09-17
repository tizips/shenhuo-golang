package shenhuo

import "github.com/tizips/shenhuo/model"

// ToArticle 文章响应
type ToArticle struct {
	ID          uint   `json:"id"`                // 文章ID
	Title       string `json:"title"`             // 文章标题
	Thumb       string `json:"thumb"`             // 缩略图地址
	Content     string `json:"content,omitempty"` // 正文内容；富文本；列表场景省略
	PublishedAt string `json:"published_at"`      // 发布时间
	IsTop       uint8  `json:"is_top"`            // 是否置顶；枚举：1=是，2=否
	IsRecommend uint8  `json:"is_recommend"`      // 是否首页推荐；枚举：1=是，2=否
}

// ToPage 单页响应
type ToPage struct {
	ID      uint   `json:"id"`                // 单页ID
	Title   string `json:"title"`             // 单页标题
	Content string `json:"content,omitempty"` // 单页内容；富文本；列表场景省略
}

// ToNav 导航响应
type ToNav struct {
	ID    uint   `json:"id"`    // 导航ID
	Title string `json:"title"` // 导航标题
	Icon  string `json:"icon"`  // 图标地址
	Type  string `json:"type"`  // 导航类型；枚举：link=外链，page=单页
	Value string `json:"value"` // 导航值；link 为跳转地址，page 为单页ID
	Order uint8  `json:"order"` // 排序序号；数值越小越靠前
}

// ToBanner 轮播图响应
type ToBanner struct {
	ID    uint   `json:"id"`    // 轮播图ID
	Title string `json:"title"` // 轮播图标题
	Image string `json:"image"` // 图片地址
	Link  string `json:"link"`  // 点击跳转链接；可为空
	Order uint8  `json:"order"` // 排序序号；数值越小越靠前
}

// ToScheduleItem 日程子项目响应
type ToScheduleItem struct {
	Name string `json:"name"` // 子项名称
	Time string `json:"time"` // 子项时间；展示用文本
}

// ToSchedule 日程响应
type ToSchedule struct {
	ID          uint             `json:"id"`          // 日程ID
	Title       string           `json:"title"`       // 日程标题
	Subtitle    string           `json:"subtitle"`    // 副标题
	Description string           `json:"description"` // 日程描述
	Time        string           `json:"time"`        // 日程时间；展示用文本
	Items       []ToScheduleItem `json:"items"`       // 子项目（一天多场安排）
	Order       uint8            `json:"order"`       // 排序序号；数值越小越靠前
}

// ToScheduleCategory 日程分类响应
type ToScheduleCategory struct {
	ID        uint         `json:"id"`        // 分类ID
	Name      string       `json:"name"`      // 分类名称
	Order     uint8        `json:"order"`     // 排序序号；数值越小越靠前
	Schedules []ToSchedule `json:"schedules"` // 分类下的日程列表
}

// ToDrawCategory 抽签类别响应
type ToDrawCategory struct {
	ID      uint     `json:"id"`      // 类别ID
	Name    string   `json:"name"`    // 类别名称
	Icon    string   `json:"icon"`    // 图标地址
	Quota   uint     `json:"quota"`   // 中签名额人数
	Status  uint8    `json:"status"`  // 抽签进度；枚举：1=未开始，2=已结束
	Order   uint8    `json:"order"`   // 排序序号；数值越小越靠前
	Drawn   int      `json:"drawn"`   // 已中签人数
	Winners []ToDraw `json:"winners"` // 中签人员列表
}

// ToDraw 中签人员响应
type ToDraw struct {
	ID        uint   `json:"id"`         // 抽签结果ID
	PersonID  string `json:"person_id"`  // 参赛人员ID
	Name      string `json:"name"`       // 人员姓名
	Number    string `json:"number"`     // 参赛号
	Unit      string `json:"unit"`       // 所属单位
	GroupName string `json:"group_name"` // 所属小组名称
	CreatedAt string `json:"created_at"` // 中签时间
}

// ToPerson 人员响应
type ToPerson struct {
	Name string `json:"name"` // 人员姓名
	Unit string `json:"unit"` // 所属公司/单位
}

// ToPersonGroup 参赛人员分组摘要响应

type ToPersonGroup struct {
	ID    uint   `json:"id"`    // 分组ID（抽签类别ID）
	Name  string `json:"name"`  // 分组名称
	Icon  string `json:"icon"`  // 图标地址
	Order uint8  `json:"order"` // 排序序号；数值越小越靠前
	Total int    `json:"total"` // 参赛人员数量
}

// ToPersonGroupOfPersons 分组参赛人员详情响应

type ToPersonGroupOfPersons struct {
	ID      uint              `json:"id"`      // 分组ID（抽签类别ID）
	Name    string            `json:"name"`    // 分组名称
	Icon    string            `json:"icon"`    // 图标地址
	Total   int               `json:"total"`   // 参赛人员数量
	Persons []ToGroupedPerson `json:"persons"` // 参赛人员列表
}

// ToGroupedPerson 分组参赛人员响应
type ToGroupedPerson struct {
	ID        string `json:"id"`         // 人员ID
	Name      string `json:"name"`       // 人员姓名
	Number    string `json:"number"`     // 参赛号
	Unit      string `json:"unit"`       // 所属单位
	GroupName string `json:"group_name"` // 所属小组名称
}

// ToScene 场景响应
type ToScene struct {
	ID     uint      `json:"id"`     // 场景ID
	Name   string    `json:"name"`   // 场景名称
	Order  uint8     `json:"order"`  // 排序序号；数值越小越靠前
	Medias []ToMedia `json:"medias"` // 场景下的媒体列表
}

// ToMedia 媒体响应
type ToMedia struct {
	ID      uint   `json:"id"`       // 媒体ID
	SceneID uint   `json:"scene_id"` // 所属场景ID
	Type    string `json:"type"`     // 媒体类型；枚举：image=图片，video=视频
	Title   string `json:"title"`    // 媒体标题
	URL     string `json:"url"`      // 媒体资源地址
	IsTop   uint8  `json:"is_top"`   // 是否置顶；枚举：1=是，2=否
}

// ToScore 成绩响应
type ToScore struct {
	ID        uint                `json:"id"`         // 成绩ID
	PersonID  string              `json:"person_id"`  // 参赛人员ID
	Name      string              `json:"name"`       // 人员姓名
	Number    string              `json:"number"`     // 参赛号
	Unit      string              `json:"unit"`       // 所属单位
	GroupName string              `json:"group_name"` // 所属小组名称
	Total     string              `json:"total"`      // 总分
	Items     []model.ShScoreItem `json:"items"`      // 各小项成绩列表
	Order     uint8               `json:"order"`      // 名次；按总分自动计算，数值越小名次越靠前
}
