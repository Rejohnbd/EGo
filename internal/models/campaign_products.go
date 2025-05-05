package models

import (
	"time"
)

type CampaignProduct struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID     uint64     `gorm:"not null;column:product_id" json:"product_id"`
	CampaignID    *uint64    `gorm:"column:campaign_id" json:"campaign_id,omitempty"`
	CampaignPrice *float64   `gorm:"type:decimal(8,2);column:campaign_price" json:"campaign_price,omitempty"`
	UnitsForSale  *int       `gorm:"column:units_for_sale" json:"units_for_sale,omitempty"`
	StartDate     *time.Time `gorm:"column:start_date" json:"start_date,omitempty"`
	EndDate       *time.Time `gorm:"column:end_date" json:"end_date,omitempty"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	// Optional relations
	Product *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	// Campaign *Campaign  `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
}

func (CampaignProduct) TableName() string {
	return "campaign_products"
}
