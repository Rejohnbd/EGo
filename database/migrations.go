package database

import (
	"dakbazar/internal/models"
	"log"
)

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Zone{},
		&models.Status{},
		&models.Admin{},
		&models.AdminRole{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed")
}
