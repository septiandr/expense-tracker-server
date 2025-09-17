package seeders

import (
	"expense-tracker/config"
	"expense-tracker/models"
	"log"
	"time"
)

func Seed() {
	// === Seed Categories ===
	var categoryCount int64
	config.DB.Model(&models.Category{}).Count(&categoryCount)

	if categoryCount == 0 {
		categories := []models.Category{
			{Name: "Makanan"},
			{Name: "Transportasi"},
			{Name: "Hiburan"},
		}

		if err := config.DB.Create(&categories).Error; err != nil {
			log.Fatalf("Gagal seed categories: %v", err)
		}
		log.Println("✅ Categories seeded")
	} else {
		log.Println("ℹ️ Categories sudah ada, skip seeding")
	}

	// === Seed Transactions ===
	var transactionCount int64
	config.DB.Model(&models.Transaction{}).Count(&transactionCount)

	if transactionCount == 0 {
		transactions := []models.Transaction{
			{CategoryID: 1, Amount: 50000, Type: "expense", Description: "Makan siang", CreatedAt: time.Now()},
			{CategoryID: 2, Amount: 20000, Type: "expense", Description: "Naik ojek", CreatedAt: time.Now()},
			{CategoryID: 3, Amount: 100000, Type: "income", Description: "Bonus kerja", CreatedAt: time.Now()},
		}

		if err := config.DB.Create(&transactions).Error; err != nil {
			log.Fatalf("Gagal seed transactions: %v", err)
		}
		log.Println("✅ Transactions seeded")
	} else {
		log.Println("ℹ️ Transactions sudah ada, skip seeding")
	}
}
