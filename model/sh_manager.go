package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShManager = "sh_manager"

type ShManager struct {
	ID        string         `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Mobile    string         `gorm:"column:mobile"`
	Password  string         `gorm:"column:password"`
	IsEnable  uint8          `gorm:"column:is_enable"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShManager) TableName() string {
	return TableShManager
}
