package models

import (
	"time"

	"gorm.io/gorm"
)

type TaxClass struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name      string         `gorm:"type:varchar(191);not null;column:name" json:"name"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`

	Options []TaxClassOption `gorm:"foreignKey:ClassID" json:"options"`
}

// TableName sets the table name correctly
func (TaxClass) TableName() string {
	return "tax_classes"
}
