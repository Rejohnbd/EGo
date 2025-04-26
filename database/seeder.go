package database

import (
	"dakbazar/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func SeedStatus() error {
	fmt.Println("Seeding Status...")

	var count int64
	if err := DB.Model(&models.Status{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("Statuses table already seeded, skipping...")
		return nil
	}

	// ➔ Read JSON file
	fileData, err := ioutil.ReadFile("data/statuses.json")
	if err != nil {
		log.Println("Failed to read statuses.json:", err)
		return err
	}

	var statuses []models.Status
	if err := json.Unmarshal(fileData, &statuses); err != nil {
		log.Println("Failed to unmarshal statuses JSON:", err)
		return err
	}

	for i := range statuses {
		err := DB.Create(&statuses[i]).Error
		if err != nil {
			log.Println("Failed to insert status:", err)
			return err
		}
	}

	fmt.Println("Statuses seeded successfully!")
	return nil
}

func SeedAdmins() error {
	fmt.Println("Seeding Admins...")

	var count int64
	if err := DB.Model(&models.Admin{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("Admins table already seeded, skipping...")
		return nil
	}

	fileData, err := ioutil.ReadFile("data/admins.json")
	if err != nil {
		log.Println("Failed to read admins.json:", err)
		return err
	}

	var admins []models.Admin
	if err := json.Unmarshal(fileData, &admins); err != nil {
		log.Println("Failed to unmarshal admins JSON:", err)
		return err
	}

	for i := range admins {
		err := DB.Create(&admins[i]).Error
		if err != nil {
			log.Println("Failed to insert admin:", err)
			return err
		}
	}

	fmt.Println("Admins seeded successfully!")
	return nil
}

func SeedZones() error {
	fmt.Println("Seeding Zones...")

	var count int64
	if err := DB.Model(&models.Zone{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("Zones table already seeded, skipping...")
		return nil
	}

	fileData, err := ioutil.ReadFile("data/zones.json")
	if err != nil {
		log.Println("Failed to read zones.json:", err)
		return err
	}

	var zones []models.Zone
	if err := json.Unmarshal(fileData, &zones); err != nil {
		log.Println("Failed to unmarshal zones JSON:", err)
		return err
	}

	for i := range zones {
		err := DB.Create(&zones[i]).Error
		if err != nil {
			log.Println("Failed to insert zone:", err)
			return err
		}
	}

	fmt.Println("Zones seeded successfully!")
	return nil
}

func SeedMediaUploads() error {
	fmt.Println("Seeding MediaUploads...")

	var count int64
	if err := DB.Model(&models.MediaUpload{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		fmt.Println("MediaUploads table already seeded, skipping...")
		return nil
	}

	fileData, err := ioutil.ReadFile("data/media_uploads.json")
	if err != nil {
		log.Println("Failed to read media_uploads.json:", err)
		return err
	}

	type MediaUploadJSON struct {
		ID         string  `json:"id"`
		Title      string  `json:"title"`
		Path       string  `json:"path"`
		Alt        *string `json:"alt"`
		Size       *string `json:"size"`
		Dimensions *string `json:"dimensions"`
		VendorID   *string `json:"vendor_id"`
		UserID     *string `json:"user_id"`
		CreatedAt  string  `json:"created_at"`
		UpdatedAt  string  `json:"updated_at"`
	}

	var mediaUploadsJSON []MediaUploadJSON
	if err := json.Unmarshal(fileData, &mediaUploadsJSON); err != nil {
		log.Println("Failed to unmarshal media_uploads JSON:", err)
		return err
	}

	for _, m := range mediaUploadsJSON {
		idUint, _ := strconv.ParseUint(m.ID, 10, 64)

		var vendorID *uint
		if m.VendorID != nil && *m.VendorID != "" {
			vendorUint, _ := strconv.ParseUint(*m.VendorID, 10, 64)
			v := uint(vendorUint)
			vendorID = &v
		}

		var userID *uint
		if m.UserID != nil && *m.UserID != "" {
			userUint, _ := strconv.ParseUint(*m.UserID, 10, 64)
			u := uint(userUint)
			userID = &u
		}

		// Parse created_at and updated_at
		var createdAt, updatedAt *time.Time
		if m.CreatedAt != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", m.CreatedAt)
			createdAt = &t
		}
		if m.UpdatedAt != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", m.UpdatedAt)
			updatedAt = &t
		}

		mediaUpload := models.MediaUpload{
			ID:         uint(idUint),
			Title:      m.Title,
			Path:       m.Path,
			Alt:        m.Alt,
			Size:       m.Size,
			Dimensions: m.Dimensions,
			VendorID:   vendorID,
			UserID:     userID,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		}

		err := DB.Create(&mediaUpload).Error
		if err != nil {
			log.Println("Failed to insert media_upload:", err)
			return err
		}
	}

	fmt.Println("MediaUploads seeded successfully!")
	return nil
}

func SeedAdminShippingMethods() error {
	fmt.Println("Seeding AdminShippingMethods...")

	var count int64
	if err := DB.Model(&models.AdminShippingMethod{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("AdminShippingMethods table already seeded, skipping...")
		return nil
	}

	fileData, err := ioutil.ReadFile("data/admin_shipping_methods.json")
	if err != nil {
		log.Println("Failed to read admin_shipping_methods.json:", err)
		return err
	}

	type AdminShippingMethodJSON struct {
		ID        string `json:"id"`
		ZoneID    string `json:"zone_id"`
		Title     string `json:"title"`
		Cost      string `json:"cost"`
		StatusID  string `json:"status_id"`
		IsDefault string `json:"is_default"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var adminShippingMethodsJSON []AdminShippingMethodJSON
	if err := json.Unmarshal(fileData, &adminShippingMethodsJSON); err != nil {
		log.Println("Failed to unmarshal admin_shipping_methods JSON:", err)
		return err
	}

	for _, a := range adminShippingMethodsJSON {
		// String to uint or int conversions
		id, _ := strconv.ParseUint(a.ID, 10, 64)
		zoneID, _ := strconv.ParseUint(a.ZoneID, 10, 64)
		statusID, _ := strconv.ParseUint(a.StatusID, 10, 64)
		isDefault, _ := strconv.ParseInt(a.IsDefault, 10, 16)
		cost, _ := strconv.ParseFloat(a.Cost, 64)

		// Parse created_at and updated_at manually
		var createdAt, updatedAt *time.Time
		if a.CreatedAt != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", a.CreatedAt)
			createdAt = &t
		}
		if a.UpdatedAt != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", a.UpdatedAt)
			updatedAt = &t
		}

		adminShippingMethod := models.AdminShippingMethod{
			ID:        uint(id),
			ZoneID:    uint(zoneID),
			Title:     a.Title,
			Cost:      cost,
			StatusID:  uint(statusID),
			IsDefault: int16(isDefault),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		err := DB.Create(&adminShippingMethod).Error
		if err != nil {
			log.Println("Failed to insert admin_shipping_method:", err)
			return err
		}
	}

	fmt.Println("AdminShippingMethods seeded successfully!")
	return nil
}
