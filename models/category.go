package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:50"`
	UserID    uint
	CreatedAt time.Time

	Transactions []Transaction `json:"transactions"`

}
