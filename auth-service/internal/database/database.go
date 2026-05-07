package database

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = database

	log.Println("Database connected successfully")

	DB.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.RefreshToken{},
		&models.Invitation{},
		&models.AuditLog{},
		&models.Permission{},
		&models.UserPermission{},
		&models.APIKey{},
	)
}