package models

import (
	"time"

	"gorm.io/gorm"
)

type Badge struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id"`
	Name      *string        `gorm:"size:191;column:name"`
	Image     *uint          `gorm:"column:image"`
	For       *string        `gorm:"size:191;column:for"`
	SaleCount *int64         `gorm:"column:sale_count"`
	Type      *int8          `gorm:"column:type"`
	Status    string         `gorm:"size:191;not null;default:'draft';column:status"`
	CreatedAt *time.Time     `gorm:"column:created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
