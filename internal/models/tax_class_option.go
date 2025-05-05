package models

import "time"

type TaxClassOption struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ClassID    uint64     `gorm:"not null;index;column:class_id" json:"class_id"`
	TaxName    string     `gorm:"type:varchar(191);not null;column:tax_name" json:"tax_name"`
	CountryID  *uint64    `gorm:"index;column:country_id" json:"country_id,omitempty"`
	StateID    *uint64    `gorm:"index;column:state_id" json:"state_id,omitempty"`
	CityID     *uint64    `gorm:"index;column:city_id" json:"city_id,omitempty"`
	PostalCode *string    `gorm:"type:varchar(191);column:postal_code" json:"postal_code,omitempty"`
	Priority   int        `gorm:"not null;column:priority" json:"priority"`
	IsCompound *bool      `gorm:"column:is_compound" json:"is_compound,omitempty"`
	IsShipping *bool      `gorm:"column:is_shipping" json:"is_shipping,omitempty"`
	Rate       float64    `gorm:"type:double(8,2);not null;column:rate" json:"rate"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (TaxClassOption) TableName() string {
	return "tax_class_options"
}
