package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name        string         `gorm:"size:191;not null;column:name" json:"name"`
	Slug        string         `gorm:"size:191;not null;column:slug" json:"slug"`
	Description *string        `gorm:"type:text;column:description" json:"description"`
	ImageID     *uint64        `gorm:"index;column:image_id" json:"image_id"`
	StatusID    *uint64        `gorm:"index;column:status_id" json:"status_id"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}
