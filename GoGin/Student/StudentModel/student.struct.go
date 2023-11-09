package studentmodel

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID         uint   `json:"user_id" gorm:"primarykey"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email" gorm:"uniqueIndex;not null"`
	Country    string `json:"country" gorm:"not null"`
	RollNumber string `json:"roll_number" gorm:"not null;uniqueIndex"`
	Age        int    `json:"age" gorm:"not null;size:3"`
	Password   string `json:"password" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
