package model

import (
	"github.com/dromara/carbon/v2"
	gocarbon "github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShBanner = "sh_banner"

type ShBanner struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Title     string          `gorm:"column:title"`
	Image     string          `gorm:"column:image"`
	Link      string          `gorm:"column:link"`
	Order     uint8           `gorm:"column:order"`
	StartedAt carbon.Carbon   `gorm:"column:started_at"` // 生效开始时间
	EndedAt   carbon.Carbon   `gorm:"column:ended_at"`   // 生效结束时间
	CreatedAt gocarbon.Carbon `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt gocarbon.Carbon `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ShBanner) TableName() string {
	return TableShBanner
}
