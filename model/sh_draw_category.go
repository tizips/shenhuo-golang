package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShDrawCategory = "sh_draw_category"

type ShDrawCategory struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Icon      string         `gorm:"column:icon"`
	Quota     uint           `gorm:"column:quota"`
	Order     uint8          `gorm:"column:order"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShDrawCategory) TableName() string {
	return TableShDrawCategory
}
