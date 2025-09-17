package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:100"`
	Email string `gorm:"uniqueIndex; size:100"`
	PasswordHash string 
	CreatedAt time.Time
}