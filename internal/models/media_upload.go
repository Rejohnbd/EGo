package models

import "time"

type MediaUpload struct {
	ID         uint       `gorm:"primaryKey;autoIncrement;index" json:"id"`
	Title      string     `gorm:"type:text;not null" json:"title"`
	Path       string     `gorm:"type:text;not null" json:"path"`
	Alt        *string    `gorm:"type:text" json:"alt,omitempty"`
	Size       *string    `gorm:"type:text" json:"size,omitempty"`
	Dimensions *string    `gorm:"type:text" json:"dimensions,omitempty"`
	VendorID   *uint      `gorm:"column:vendor_id;index" json:"vendor_id,omitempty"`
	UserID     *uint      `gorm:"column:user_id;index" json:"user_id,omitempty"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}
