package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShAppeal = "sh_appeal"

type ShAppeal struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	PersonID  string         `gorm:"column:person_id"`
	Reason    string         `gorm:"column:reason"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Person *ShPerson `gorm:"foreignKey:PersonID;references:ID"`
}

func (ShAppeal) TableName() string {
	return TableShAppeal
}
