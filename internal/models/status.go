package models

import "gorm.io/gorm"

type Status struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
