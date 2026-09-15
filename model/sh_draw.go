package model

import (
	"github.com/golang-module/carbon/v2"
)

const TableShDraw = "sh_draw"

type ShDraw struct {
	ID         uint          `gorm:"column:id;primaryKey"`
	CategoryID uint          `gorm:"column:category_id"`
	PersonID   string        `gorm:"column:person_id"`
	CreatedAt  carbon.Carbon `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt  carbon.Carbon `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`

	Category *ShDrawCategory `gorm:"foreignKey:CategoryID;references:ID"`
	Person   *ShPerson       `gorm:"foreignKey:PersonID;references:ID"`
}

func (ShDraw) TableName() string {
	return TableShDraw
}
