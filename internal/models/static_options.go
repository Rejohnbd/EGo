package models

import "time"

type StaticOption struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	OptionName  string     `gorm:"type:varchar(191);not null;uniqueIndex;column:option_name" json:"option_name"`
	OptionValue *string    `gorm:"type:longtext;column:option_value" json:"option_value,omitempty"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// TableName explicitly sets the table name
func (StaticOption) TableName() string {
	return "static_options"
}
