package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type NewUserRequest struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Email     string `json:"email" `
	Country   string `json:"country" `
	Role      string `json:"role" `
	Age       int    `json:"age"`
	Password  string `json:"password" `
}
type LoginUserRequest struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type GetUserRequest struct {
	ID uint `json:"user_id"`
}

type UpdateUserRequest struct {
	ID        uint   `json:"user_id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Email     string `json:"email" `
	Country   string `json:"country" `
	Role      string `json:"role" `
	Age       int    `json:"age"`
	Password  string `json:"password"`
}

/*
I have used the

Try ozzo-validation https://github.com/go-ozzo/ozzo-validation

for validation of the schema of body data which passed in form of json type
*/

func (a NewUserRequest) NewUserValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.FirstName, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.LastName, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.Email, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
		validation.Field(&a.Role, validation.Required),
		validation.Field(&a.Age, validation.Required),
		validation.Field(&a.Country, validation.Required),
		validation.Field(&a.Password, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]{6}$`))),
	)
}

func (a LoginUserRequest) LoginValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)).Error("Please Enter Valid Email Adress")),
		validation.Field(&a.Password, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$")).Error("Password must be in 6 digit")),
	)
}

func (a GetUserRequest) UserIdValidate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ID, validation.Required))
}
