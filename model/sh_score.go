package model

import (
	"encoding/json"

	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShScore = "sh_score"

type ShScore struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	PersonID  string         `gorm:"column:person_id"`
	Total     string         `gorm:"column:total"`
	Items     string         `gorm:"column:items"`
	Order     uint8          `gorm:"column:order"` // 名次；按总分自动计算，数值越小名次越靠前
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Person *ShPerson `gorm:"foreignKey:PersonID;references:ID"`
}

func (ShScore) TableName() string {
	return TableShScore
}

type ShScoreItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s ShScore) ItemList() []ShScoreItem {
	if s.Items == "" {
		return []ShScoreItem{}
	}

	var items []ShScoreItem
	if err := json.Unmarshal([]byte(s.Items), &items); err != nil {
		return []ShScoreItem{}
	}

	return items
}

func MarshalScoreItems(items []ShScoreItem) string {
	if len(items) == 0 {
		return "[]"
	}

	raw, err := json.Marshal(items)
	if err != nil {
		return "[]"
	}

	return string(raw)
}
