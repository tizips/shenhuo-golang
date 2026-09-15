package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShBanner = "sh_banner"

type ShBanner struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Title     string         `gorm:"column:title"`
	Image     string         `gorm:"column:image"`
	Link      string         `gorm:"column:link"`
	Order     uint8          `gorm:"column:order"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShBanner) TableName() string {
	return TableShBanner
}
