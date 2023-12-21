package main

import (
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.DBInitializer()
}

func main() {
	initializer.DB.AutoMigrate(&models.User{})
	initializer.DB.AutoMigrate(&models.Product{})
	initializer.DB.AutoMigrate(&models.Order{})
	initializer.DB.AutoMigrate(&models.OrderDetail{})
}
