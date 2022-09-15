package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"size:255" json:"name"`
	Email       string `gorm:"size:255" json:"email"`
	Address     string `gorm:"size:255" json:"address"`
	Gender      string `gorm:"size:255" json:"gender"`
	PhoneNumber string `gorm:"size:255" json:"phone_number"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
