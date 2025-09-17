package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// TODO: Tambah middleware JWT
	// r.GET("/categories", controllers.GetCategories)
	// r.POST("/transactions", controllers.CreateTransaction)
}
