package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShPerson = "sh_person"

type ShPerson struct {
	ID                 string         `gorm:"column:id;primaryKey"`
	Name               string         `gorm:"column:name"`
	Unit               string         `gorm:"column:unit"`
	Mobile             string         `gorm:"column:mobile"`
	Password           string         `gorm:"column:password"`
	Number             string         `gorm:"column:number"`
	GroupName          string         `gorm:"column:group_name"`
	OpenID             string         `gorm:"column:openid"`
	MustChangePassword uint8          `gorm:"column:must_change_password"`
	CreatedAt          carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt          carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ShPerson) TableName() string {
	return TableShPerson
}
