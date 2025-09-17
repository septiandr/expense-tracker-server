package controllers

import (
	"expense-tracker/middleware"
	"expense-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	db := middleware.GetDB(c) // ambil koneksi DB dari context

	var transactions []models.Transaction
	if err := db.Preload("Category").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
