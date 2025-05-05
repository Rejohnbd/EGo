package models

import (
	"time"

	"gorm.io/gorm"
)

type VendorAddress struct {
	ID                    uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	VendorID              uint64         `gorm:"index;not null;column:vendor_id" json:"vendor_id"`
	CountryID             *uint64        `gorm:"index;column:country_id" json:"country_id"`
	StateID               *uint64        `gorm:"index;column:state_id" json:"state_id"`
	CityID                *uint64        `gorm:"index;column:city_id" json:"city_id"`
	ZipCode               *string        `gorm:"size:191;column:zip_code" json:"zip_code"`
	Address               *string        `gorm:"type:text;column:address" json:"address"`
	DivisionID            *uint64        `gorm:"column:division_id" json:"division_id"`
	DistrictID            *uint64        `gorm:"column:district_id" json:"district_id"`
	PoliceStationID       *uint64        `gorm:"column:police_station_id" json:"police_station_id"`
	PostOfficeID          *uint64        `gorm:"column:post_office_id" json:"post_office_id"`
	PickupDivisionID      *uint64        `gorm:"column:pickup_division_id" json:"pickup_division_id"`
	PickupCountryID       *uint64        `gorm:"column:pickup_country_id" json:"pickup_country_id"`
	PickupDistrictID      *uint64        `gorm:"column:pickup_district_id" json:"pickup_district_id"`
	PickupPoliceStationID *uint64        `gorm:"column:pickup_police_station_id" json:"pickup_police_station_id"`
	PickupPostOfficeID    *uint64        `gorm:"column:pickup_post_office_id" json:"pickup_post_office_id"`
	PickupPostCode        *string        `gorm:"size:191;column:pickup_post_code" json:"pickup_post_code"`
	PickupAddress         *string        `gorm:"size:191;column:pickup_address" json:"pickup_address"`
	CreatedAt             *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`
}
