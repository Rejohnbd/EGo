package models

import (
	"time"

	"gorm.io/gorm"
)

type ChildCategory struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CategoryID    uint64         `gorm:"not null;index;column:category_id" json:"category_id"`
	SubCategoryID uint64         `gorm:"not null;index;column:sub_category_id" json:"sub_category_id"`
	Name          string         `gorm:"type:varchar(191);not null;column:name" json:"name"`
	Slug          string         `gorm:"type:varchar(191);not null;column:slug" json:"slug"`
	Description   *string        `gorm:"type:tinytext;column:description" json:"description,omitempty"`
	ImageID       *uint64        `gorm:"index;column:image_id" json:"image_id,omitempty"`
	StatusID      *uint64        `gorm:"index;column:status_id" json:"status_id,omitempty"`
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`

	// Optional Relations
	Category    *Category    `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	SubCategory *SubCategory `gorm:"foreignKey:SubCategoryID" json:"sub_category,omitempty"`
}

func (ChildCategory) TableName() string {
	return "child_categories"
}
