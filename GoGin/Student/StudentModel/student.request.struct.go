package studentmodel

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetStudentRequest struct {
	RollNumber string `json:"roll_number"`
}

type CreateStudentRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email" gorm:"uniqueIndex;not null"`
	Country    string `json:"country" gorm:"not null"`
	RollNumber string `json:"roll_number" gorm:"not null;uniqueIndex"`
	Age        int    `json:"age" gorm:"not null;size:3"`
	ClassID    uint   `json:"class_id"`
	Password   string `json:"password" gorm:"not null"`
}
type UpdateStudentRequest struct {
	ID         uint   `json:"student_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email" gorm:"uniqueIndex;not null"`
	Country    string `json:"country" gorm:"not null"`
	RollNumber string `json:"roll_number" gorm:"not null;uniqueIndex"`
	Age        int    `json:"age" gorm:"not null;size:3"`
	ClassID    uint   `json:"class_id"`
	Password   string `json:"password" gorm:"not null"`
}

type GetClassWiseStudentReq struct {
	ClassID uint `json:"class_id"`
}

type OutputClassWiseStudent struct {
	FirstName string `json:"first_name"`
	ClassName string `json:"class_name"`
}

func (a CreateStudentRequest) NewStudentValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.FirstName, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.LastName, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
		validation.Field(&a.RollNumber, validation.Required),
		validation.Field(&a.Age, validation.Required),
		validation.Field(&a.ClassID, validation.Required),
		validation.Field(&a.Country, validation.Required),
		validation.Field(&a.Password, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]{6}$`))),
	)
}

func (a GetClassWiseStudentReq) GetClassWiseStudentValidate() error {
	return validation.ValidateStruct(&a, validation.Field(&a.ClassID, validation.Required))
}
