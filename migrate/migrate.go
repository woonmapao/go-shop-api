package main

import (
	"log"

	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

func init() {
	// Initialize environment variables and database
	initializer.LoadEnvVariables()
	initializer.DBInitializer()
}

func main() {
	// Perform database auto migrations

	// Specify models to be migrated
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderDetail{},
	}

	// Auto migrate the specified models
	err := initializer.DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Fatalf("Failed to perform auto migration: %v", err)
	}
}
