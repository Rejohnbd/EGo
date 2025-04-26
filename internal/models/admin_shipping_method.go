package models

import (
	"time"
)

type AdminShippingMethod struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ZoneID    uint       `gorm:"column:zone_id;not null" json:"zone_id"`
	Title     string     `gorm:"type:varchar(191);not null" json:"title"`
	Cost      float64    `gorm:"type:decimal(8,2);not null" json:"cost"`
	StatusID  uint       `gorm:"column:status_id;not null" json:"status_id"`
	IsDefault int16      `gorm:"type:smallint;default:0" json:"is_default"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}
