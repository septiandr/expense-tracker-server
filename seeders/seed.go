package seeders

import (
	"expense-tracker/config"
	"expense-tracker/models"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// Hash password
	password, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)

	// Seed user
	user := models.User{
		Name:         "Andi",
		Email:        "andi@mail.com",
		PasswordHash: string(password),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("⚠️ User seeding error:", err)
	}

	// Seed categories
	categories := []models.Category{
		{Name: "Makanan", UserID: user.ID},
		{Name: "Transportasi", UserID: user.ID},
		{Name: "Hiburan", UserID: user.ID},
	}
	for _, cat := range categories {
		config.DB.Create(&cat)
	}

	// Seed transactions
	transactions := []models.Transaction{
		{UserID: user.ID, CategoryID: 1, Amount: 50000, Type: "expense", Description: "Makan siang", CreatedAt: time.Now()},
		{UserID: user.ID, CategoryID: 2, Amount: 20000, Type: "expense", Description: "Naik ojek", CreatedAt: time.Now()},
		{UserID: user.ID, CategoryID: 3, Amount: 100000, Type: "income", Description: "Bonus kerja", CreatedAt: time.Now()},
	}
	for _, trx := range transactions {
		config.DB.Create(&trx)
	}

	fmt.Println("✅ Database seeded successfully")
}
