package controllers

import (
	"expense-tracker/middleware"
	"expense-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	db := middleware.GetDB(c) // ambil koneksi DB dari context

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func CreateCategory(c *gin.Context) {
	db := middleware.GetDB(c)

	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{Name: input.Name}
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}
