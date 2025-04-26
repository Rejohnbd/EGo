package models

import "time"

type Zone struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;index" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
