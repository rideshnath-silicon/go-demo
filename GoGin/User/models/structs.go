package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `json:"user_id" gorm:"primarykey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	PhoneNumber string `json:"phone_number"`
	Country     string `json:"country" gorm:"not null"`
	Role        string `json:"role" gorm:"not null"`
	Age         int    `json:"age" gorm:"not null;size:3"`
	Password    string `json:"password" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
