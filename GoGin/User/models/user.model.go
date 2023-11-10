package models

import (
	"gin/config"
	"gin/helpers"
	"time"
)

func LoginUser(email string, pass string) (User, error) {
	var user User
	result := config.DB.Select("email", "id").Where("email = ? ", email).Where("password = ?", pass).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetAllUser() []User {
	var user []User
	result := config.DB.Find(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}

func GetUserByEmail(Email string) (string, error) {
	var user User
	result := config.DB.Select("password").Where("email = ?", Email).Find(&user)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", nil
	}
	return user.Password, nil
}

func GetUser(id interface{}) interface{} {
	var user User
	result := config.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return 0
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return user
}
func NewUser(data NewUserRequest) interface{} {
	pass, err := helpers.HashData(data.Password)
	if err != nil {
		return err.Error()
	}
	user := User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Age:       data.Age,
		Email:     data.Email,
		Country:   data.Country,
		Password:  pass,
		Role:      data.Role,
		CreatedAt: time.Now(),
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error.Error()
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return user
}

func UpdateUser(id uint, data UpdateUserRequest) interface{} {
	pass, err := helpers.HashData(data.Password)
	if err != nil {
		return err.Error()
	}
	var user = User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  pass,
		Role:      data.Role,
		Country:   data.Country,
		Age:       data.Age,
		UpdatedAt: time.Now(),
	}
	result := config.DB.Where("id = ?", id).Updates(user).Model(&user)
	if result.Error != nil {
		return result.Error.Error()
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return data
}

func DeleteUser(id uint) (interface{}, error) {
	var user User
	err := config.DB.Delete(&user).Where("id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected == 0 {
		return 0, nil
	}
	return user, nil
}
