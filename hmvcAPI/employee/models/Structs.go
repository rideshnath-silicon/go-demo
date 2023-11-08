package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string    `gorm:"not null"`
	Country   string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
