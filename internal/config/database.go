package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hospital-middleware/internal/model"
	"log"
)

func ConnectDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Patient{},
		&model.Hospital{},
		&model.PatientHospital{},
		&model.Staff{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
