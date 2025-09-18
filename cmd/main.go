package main

import (
	"expense-tracker/config"
	"expense-tracker/middleware"
	"expense-tracker/models"
	"expense-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect DB
	config.ConnectDB()

	// Auto migrate models
	config.DB.AutoMigrate(&models.Category{}, &models.Transaction{})

	// seeders.Seed()
	r.Use(middleware.DBMiddleware())

	// Setup routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}
