package classmodel

import validation "github.com/go-ozzo/ozzo-validation"

type NewClassRequest struct {
	ClassName string `json:"class_name"`
	UserId    uint   `json:"user_id"`
	SubjectId uint   `json:"subject_id"`
}
type UpdateClassRequest struct {
	ID        uint   `json:"class_id"`
	ClassName string `json:"class_name"`
	UserId    uint   `json:"user_id"`
	SubjectId uint   `json:"subject_id"`
}

func (a NewClassRequest) NewClassValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ClassName, validation.Required, validation.Length(1, 10)),
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.SubjectId, validation.Required),
	)
}
