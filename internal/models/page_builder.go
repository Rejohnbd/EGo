package models

import "time"

type PageBuilder struct {
	ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	AddonName      *string    `gorm:"type:varchar(255)" json:"addon_name,omitempty"`
	AddonType      *string    `gorm:"type:varchar(255)" json:"addon_type,omitempty"`
	AddonLocation  *string    `gorm:"type:varchar(255)" json:"addon_location,omitempty"`
	AddonOrder     *uint      `gorm:"type:bigint" json:"addon_order,omitempty"`
	AddonPageID    *uint      `gorm:"type:bigint" json:"addon_page_id,omitempty"`
	AddonPageType  *string    `gorm:"type:varchar(255)" json:"addon_page_type,omitempty"`
	AddonSettings  *string    `gorm:"type:longtext" json:"addon_settings,omitempty"`
	AddonNamespace *string    `gorm:"type:text" json:"addon_namespace,omitempty"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}
