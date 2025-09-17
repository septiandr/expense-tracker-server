package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CategoryID  uint      `json:"category_id"`
	Amount      float64   `json:"amount"`
	Type        string    `gorm:"size:10" json:"type"` // income / expense
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`

	// Relasi: setiap transaksi punya 1 user & 1 kategori
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}
