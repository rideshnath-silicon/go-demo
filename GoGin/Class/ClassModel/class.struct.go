package classmodel

import "time"

type Class struct {
	ID        uint   `gorm:"primarykey"`
	ClassName string `gorm:"uniqueIndex;not null"`
	UserId    uint 
	SubjectId uint
	CreatedAt time.Time
}
