package studentmodel

import (
	"gin/config"
	"time"
)

func GetStudent(roll_number string) (interface{}, error) {
	var student Student
	result := config.DB.Where("roll_number = ?", roll_number).Find(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return student, nil
}

func RegisterStudent(data CreateStudentRequest) (interface{}, error) {
	var student = Student{
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		RollNumber: data.RollNumber,
		Age:        data.Age,
		Password:   data.Password,
		Country:    data.Country,
		ClassID:    data.ClassID,
		CreatedAt:  time.Now(),
	}
	result := config.DB.Create(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return student, nil
}

func UpdateStudent(id uint, data UpdateStudentRequest) (interface{}, error) {
	var student = Student{
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		RollNumber: data.RollNumber,
		Age:        data.Age,
		Password:   data.Password,
		Country:    data.Country,
		ClassID:    data.ClassID,
		UpdatedAt:  time.Now(),
	}
	result := config.DB.Where("id = ?", id).Updates(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return student, nil
}

func GetClassWiseStudent(id uint) (interface{}, error) {
	var student Student
	var results []OutputClassWiseStudent
	err := config.DB.Select("students.first_name", "classes.class_name").Joins("Join classes ON classes.id = students.class_id").Where("class_id", id).Find(&student).Scan(&results)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected == 0 {
		return 0, err.Error
	}
	return results, nil
}
