package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShDrawCategory = "sh_draw_category"

// 抽签进度
const (
	ShDrawCategoryStatusOfNotStarted uint8 = 1 // 未开始
	ShDrawCategoryStatusOfFinished   uint8 = 2 // 已结束
)

type ShDrawCategory struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Icon      string         `gorm:"column:icon"`
	Quota     uint           `gorm:"column:quota"`
	Status    uint8          `gorm:"column:status"` // 抽签进度；枚举：1=未开始，2=已结束
	Order     uint8          `gorm:"column:order"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShDrawCategory) TableName() string {
	return TableShDrawCategory
}
