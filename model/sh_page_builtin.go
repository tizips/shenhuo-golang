package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShPageBuiltin = "sh_page_builtin"

type ShPageBuiltin struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Key       string         `gorm:"column:key"`
	PageID    uint           `gorm:"column:page_id"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShPageBuiltin) TableName() string {
	return TableShPageBuiltin
}
