package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShArticle = "sh_article"

type ShArticle struct {
	ID          uint           `gorm:"column:id;primaryKey"`
	Title       string         `gorm:"column:title"`
	Thumb       string         `gorm:"column:thumb"`
	Content     string         `gorm:"column:content"`
	PublishedAt carbon.Carbon  `gorm:"column:published_at" carbon:"type:dateTime"`
	IsTop       uint8          `gorm:"column:is_top"`
	IsRecommend uint8          `gorm:"column:is_recommend"`
	CreatedAt   carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt   carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShArticle) TableName() string {
	return TableShArticle
}
