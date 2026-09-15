package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShMedia = "sh_media"

type ShMedia struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	SceneID   uint           `gorm:"column:scene_id"`
	Type      string         `gorm:"column:type"`
	Title     string         `gorm:"column:title"`
	URL       string         `gorm:"column:url"`
	IsTop     uint8          `gorm:"column:is_top"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Scene *ShScene `gorm:"foreignKey:SceneID;references:ID"`
}

func (ShMedia) TableName() string {
	return TableShMedia
}

const (
	ShMediaOfTypeImage = "image"
	ShMediaOfTypeVideo = "video"
)
