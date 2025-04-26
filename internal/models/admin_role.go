package models

import "time"

type AdminRole struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null;" json:"name"`
	Permissions string     `gorm:"longtext:not null;" json:"permissions"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
