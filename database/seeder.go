package database

import (
	"dakbazar/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
