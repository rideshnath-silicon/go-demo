package models

import (
	"time"

	"gorm.io/gorm"
)

func (UserJson) TableName() string {
	return "users" // Specify the actual table name
}

type User struct {
	ID        uint   `json:"user_id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Country   string `json:"country" gorm:"not null"`
	Role      string `json:"role" gorm:"not null"`
	Age       int    `json:"age" gorm:"not null;size:3"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

type UserJson struct {
	ID        uint   `json:"user_id" gorm:"-"`
	FirstName string `json:"first_name" gorm:"-"`
	LastName  string `json:"last_name" gorm:"-"`
	Email     string `json:"email" gorm:"-" `
	Country   string `json:"country" gorm:"-"`
	Role      string `json:"role" gorm:"-"`
	Age       int    `json:"age" gorm:"-"`
}
