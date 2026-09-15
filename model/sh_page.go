package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShPage = "sh_page"

type ShPage struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Title     string         `gorm:"column:title"`
	Content   string         `gorm:"column:content"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShPage) TableName() string {
	return TableShPage
}
