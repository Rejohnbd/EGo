package models

import (
	"time"
)

type CampaignSoldProduct struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID   *uint64    `gorm:"column:product_id" json:"product_id"`
	SoldCount   *int       `gorm:"column:sold_count" json:"sold_count"`
	TotalAmount *float64   `gorm:"column:total_amount" json:"total_amount"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
