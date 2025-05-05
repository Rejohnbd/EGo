package models

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name      string         `gorm:"type:varchar(191);not null;index;column:name" json:"name"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (Unit) TableName() string {
	return "units"
}
