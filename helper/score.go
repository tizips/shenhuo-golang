package helper

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/tizips/shenhuo/model"
)

// ScoreItemNames 成绩固定小项名称；顺序即存储与展示顺序。
var ScoreItemNames = []string{
	"理论考试",
	"管理能力考试",
	"井下紧急避险考核",
	"自救器操作考核",
	"心肺复苏考核",
}

var scoreSkipHeaders = map[string]struct{}{
	"姓名":   {},
	"单位":   {},
	"手机号":  {},
	"手机":   {},
	"小组名称": {},
	"小组":   {},
	"密码":   {},
	"总分":   {},
}

// ParseScoreHeader 解析成绩导入表头，返回参赛号列下标与各固定小项列下标。
// 总分列（如存在）会被忽略，总分由各小项成绩自动计算。
func ParseScoreHeader(headers []string) (numberCol int, itemCols []int, err error) {
	numberCol = -1
	itemCols = make([]int, len(ScoreItemNames))
	for i := range itemCols {
		itemCols[i] = -1
	}

	for i, raw := range headers {
		name := strings.TrimSpace(raw)
		if name == "" {
			continue
		}

		switch name {
		case "参赛号", "编号":
			if numberCol < 0 {
				numberCol = i
			}
		default:
			if _, skip := scoreSkipHeaders[name]; skip {
				continue
			}
			for j, item := range ScoreItemNames {
				if name == item && itemCols[j] < 0 {
					itemCols[j] = i
					break
				}
			}
		}
	}

	if numberCol < 0 {
		return 0, nil, fmt.Errorf("缺少参赛号列")
	}

	var missing []string
	for j, col := range itemCols {
		if col < 0 {
			missing = append(missing, ScoreItemNames[j])
		}
	}
	if len(missing) > 0 {
		return 0, nil, fmt.Errorf("缺少小项列：%s", strings.Join(missing, "、"))
	}

	return numberCol, itemCols, nil
}

// ScoreItemsFromRow 按固定小项顺序提取一行成绩。
func ScoreItemsFromRow(row []string, cols []int) []model.ShScoreItem {
	items := make([]model.ShScoreItem, 0, len(ScoreItemNames))
	for i, name := range ScoreItemNames {
		col := -1
		if i < len(cols) {
			col = cols[i]
		}
		items = append(items, model.ShScoreItem{
			Name:  name,
			Value: Cell(row, col),
		})
	}
	return items
}

// CalculateScoreTotal 根据各小项成绩计算总分；空白小项按 0 分计。
func CalculateScoreTotal(values []string) (string, error) {
	var total float64
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		score, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", fmt.Errorf("成绩“%s”不是有效数字", value)
		}
		total += score
	}
	return FormatScore(total), nil
}

// FormatScore 格式化分数；整数不带小数位。
func FormatScore(total float64) string {
	return strconv.FormatFloat(total, 'f', -1, 64)
}

// ParseScoreValue 解析分数；无法解析时返回负无穷，排序时垫底。
func ParseScoreValue(value string) float64 {
	score, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return math.Inf(-1)
	}
	return score
}

// RefreshScoreRanks 按总分从高到低重新计算全部成绩名次并写入 order 字段；总分相同名次并列。
func RefreshScoreRanks(db *gorm.DB) error {
	var scores []model.ShScore
	if err := db.Model(&model.ShScore{}).Select("`id`", "`total`", "`order`").Find(&scores).Error; err != nil {
		return err
	}

	sort.SliceStable(scores, func(i, j int) bool {
		ti, tj := ParseScoreValue(scores[i].Total), ParseScoreValue(scores[j].Total)
		if ti != tj {
			return ti > tj
		}
		return scores[i].ID < scores[j].ID
	})

	var prev float64
	prevRank := 0
	for i := range scores {
		total := ParseScoreValue(scores[i].Total)
		rank := i + 1
		if i > 0 && total == prev {
			rank = prevRank
		}
		prev, prevRank = total, rank

		if rank > 255 { // order 列为 tinyint unsigned
			rank = 255
		}
		if uint8(rank) == scores[i].Order {
			continue
		}
		if err := db.Model(&model.ShScore{}).Where("`id`=?", scores[i].ID).Update("`order`", uint8(rank)).Error; err != nil {
			return err
		}
	}

	return nil
}
