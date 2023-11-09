package submodel

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	ID         uint   `json:"user_id" gorm:"primarykey"`
	SubName    string `json:"sub_name"`
	AuthorName string `json:"author_name"`
	Number     string `json:"number" gorm:"not null;uniqueIndex"`
	CreatedAt  time.Time
	DeletedAt gorm.DeletedAt
}


