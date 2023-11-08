package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex;not null"`
	Country   string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Age       int    `gorm:"not null;size:3"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
