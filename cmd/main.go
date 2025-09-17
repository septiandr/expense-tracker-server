package main

import (
	"expense-tracker/config"
	"expense-tracker/models"
	"expense-tracker/routes"
	"expense-tracker/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect DB
	config.ConnectDB()
	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Transaction{})

	seeders.Seed()

	// Setup routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}
