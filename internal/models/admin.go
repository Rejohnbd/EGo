package models

import "time"

type Admin struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID          string     `gorm:"type:varchar(36);uniqueIndex;index" json:"uuid"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	Username      string     `gorm:"type:varchar(255);not null;uniqueIndex" json:"username"`
	Email         string     `gorm:"type:varchar(191);not null" json:"email"`
	Phone         string     `gorm:"type:varchar(191);index" json:"phone"`
	EmailVerified int        `gorm:"type:int;default:0" json:"email_verified"`
	Role          string     `gorm:"type:varchar(191);not null;default:'editor'" json:"role"`
	Image         string     `gorm:"type:varchar(191)" json:"image"`
	Password      string     `gorm:"type:varchar(191);not null" json:"password"`
	Status        string     `gorm:"type:varchar(191)" json:"status"`
	RememberToken string     `gorm:"type:varchar(100)" json:"remember_token"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
