package models

import "time"

type ProductRatings struct {
	ID        uint       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID uint       `gorm:"not null;index;column:product_id" json:"product_id"`
	UserID    uint       `gorm:"not null;index;column:user_id" json:"user_id"`
	Rating    int        `gorm:"not null;column:rating" json:"rating"`
	ReviewMsg *string    `gorm:"type:longtext;column:review_msg" json:"review_msg"`
	Status    *int16     `gorm:"type:smallint;column:status" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ProductRatings) TableName() string {
	return "product_ratings"
}
