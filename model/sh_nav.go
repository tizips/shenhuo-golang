package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShNav = "sh_nav"

type ShNav struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Title     string         `gorm:"column:title"`
	Icon      string         `gorm:"column:icon"`
	Type      string         `gorm:"column:type"`
	Value     string         `gorm:"column:value"`
	Order     uint8          `gorm:"column:order"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShNav) TableName() string {
	return TableShNav
}

const (
	ShNavOfTypeLink = "link"
	ShNavOfTypePage = "page"
)
