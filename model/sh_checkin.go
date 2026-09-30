package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShCheckin = "sh_checkin"

type ShCheckin struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShCheckin) TableName() string {
	return TableShCheckin
}
