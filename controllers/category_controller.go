package controllers

import (
	"expense-tracker/middleware"
	"expense-tracker/models"
	"expense-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	db := middleware.GetDB(c) // ambil koneksi DB dari context

	var categories []models.Category
	result, err := utils.Paginate(c, db, &models.Category{}, &categories, "name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
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

func DeleteCategory(c *gin.Context) {
	db := middleware.GetDB(c)

	id := c.Param("id")

	var category models.Category                          // init tempat penyimpanan data category
	if err := db.First(&category, id).Error; err != nil { // cari category berdasarkan id dan simpan di variable category
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := db.Delete(&category).Error; err != nil { // hapus category yang sudah ditemukan sesuai dengan yang ada di dalam variable category
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
