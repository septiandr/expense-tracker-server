package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint
	CategoryID  uint
	Amount      float64
	Type        string    `gorm:"size:10"` // income or expense
	Description string
	CreatedAt   time.Time
}
