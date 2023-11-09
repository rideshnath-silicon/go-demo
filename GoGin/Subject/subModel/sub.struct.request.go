package submodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AddSubjectRequest struct {
	SubName    string `json:"sub_name"`
	AuthorName string `json:"author_name"`
	Number     string `json:"number" gorm:"not null;uniqueIndex"`
}

type UpdateSubjectRequest struct {
	ID         uint   `json:"sub_id"`
	SubName    string `json:"sub_name"`
	AuthorName string `json:"author_name"`
	Number     string `json:"number"`
}

func (a AddSubjectRequest) NewSubValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.SubName, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.AuthorName, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Number, validation.Required),
	)
}
