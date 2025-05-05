package models

import (
	"time"
)

type AdminShopManage struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	StoreName    string     `gorm:"type:varchar(191);not null;column:store_name" json:"store_name"`
	LogoID       uint64     `gorm:"column:logo_id;not null" json:"logo_id"`
	CoverPhotoID uint64     `gorm:"column:cover_photo_id;not null" json:"cover_photo_id"`
	CountryID    uint64     `gorm:"column:country_id;not null" json:"country_id"`
	StateID      *uint64    `gorm:"column:state_id" json:"state_id"`
	City         *string    `gorm:"type:varchar(191);column:city" json:"city"`
	ZipCode      *string    `gorm:"type:varchar(191);column:zipcode" json:"zipcode"`
	Address      *string    `gorm:"type:text;column:address" json:"address"`
	Location     *string    `gorm:"type:varchar(191);column:location" json:"location"`
	Number       *string    `gorm:"type:varchar(191);column:number" json:"number"`
	Email        *string    `gorm:"type:varchar(191);column:email" json:"email"`
	FacebookURL  *string    `gorm:"type:varchar(191);column:facebook_url" json:"facebook_url"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (AdminShopManage) TableName() string {
	return "admin_shop_manages"
}
