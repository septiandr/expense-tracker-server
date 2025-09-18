package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/transactions", controllers.GetAllTransactions)

	r.GET("/categories", controllers.GetAllCategories)
	category := r.Group("/category")
	{
		category.POST("/", controllers.CreateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}
}
