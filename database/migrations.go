package database

import (
	"dakbazar/internal/models"
	"log"
)

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.MediaUpload{},
		&models.Zone{},
		&models.Status{},
		&models.Admin{},
		&models.AdminRole{},
		&models.AdminShippingMethod{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed")
}
