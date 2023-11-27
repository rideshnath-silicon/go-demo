package models

import (
	"time"

	"gorm.io/gorm"
)

type CarUser struct {
	ID           uint   `json:"user_id" gorm:"primarykey"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	MobileNumber string `json:"mobile_number" gorm:"not null"`
	Password     string `json:"password" gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
