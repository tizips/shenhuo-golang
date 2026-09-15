package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShSchedule = "sh_schedule"

// ShScheduleItem 日程子项目（一天多场安排）
type ShScheduleItem struct {
	Name string `json:"name"` // 子项名称
	Time string `json:"time"` // 子项时间；展示用文本
}

type ShSchedule struct {
	ID          uint             `gorm:"column:id;primaryKey"`
	CategoryID  uint             `gorm:"column:category_id"`
	Title       string           `gorm:"column:title"`
	Subtitle    string           `gorm:"column:subtitle"`
	Description string           `gorm:"column:description"`
	Time        string           `gorm:"column:time"`
	Items       []ShScheduleItem `gorm:"column:items;serializer:json"`
	Order       uint8            `gorm:"column:order"`
	CreatedAt   carbon.Carbon    `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt   carbon.Carbon    `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt   gorm.DeletedAt   `gorm:"column:deleted_at"`
}

func (ShSchedule) TableName() string {
	return TableShSchedule
}
