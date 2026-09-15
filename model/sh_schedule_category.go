package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShScheduleCategory = "sh_schedule_category"

type ShScheduleCategory struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Order     uint8          `gorm:"column:order"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShScheduleCategory) TableName() string {
	return TableShScheduleCategory
}
