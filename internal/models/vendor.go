package models

import "time"

type Vendor struct {
	ID                  uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID                *string    `gorm:"type:char(36);index" json:"uuid,omitempty"`
	OwnerName           *string    `gorm:"type:varchar(191)" json:"owner_name,omitempty"`
	BusinessName        *string    `gorm:"type:varchar(191)" json:"business_name,omitempty"`
	Description         *string    `gorm:"type:varchar(191)" json:"description,omitempty"`
	BusinessTypeID      *uint      `gorm:"column:business_type_id;index" json:"business_type_id,omitempty"`
	StatusID            uint       `gorm:"column:status_id;index;default:2" json:"status_id"`
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CheckOnlineStatus   *time.Time `gorm:"column:check_online_status" json:"check_online_status,omitempty"`
	Username            string     `gorm:"type:varchar(191);not null" json:"username"`
	Password            string     `gorm:"type:varchar(191);not null" json:"-"`
	RememberToken       *string    `gorm:"type:varchar(191)" json:"remember_token,omitempty"`
	ImageID             *uint      `gorm:"column:image_id;index" json:"image_id,omitempty"`
	Email               *string    `gorm:"type:varchar(191)" json:"email,omitempty"`
	EmailVerified       string     `gorm:"type:varchar(191);default:'0'" json:"email_verified"`
	EmailVerifyToken    *string    `gorm:"type:varchar(191)" json:"email_verify_token,omitempty"`
	CommissionType      *string    `gorm:"type:varchar(191)" json:"commission_type,omitempty"`
	CommissionAmount    *string    `gorm:"type:varchar(191)" json:"commission_amount,omitempty"`
	FirebaseDeviceToken *string    `gorm:"type:varchar(191)" json:"firebase_device_token,omitempty"`
	ProfileStatus       int8       `gorm:"type:tinyint;default:0" json:"profile_status"`

	Status *Status `gorm:"foreignKey:ProfileStatus" json:"status,omitempty"`
}
