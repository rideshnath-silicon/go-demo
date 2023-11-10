package classmodel

import (
	"gin/config"
	"time"
)

func NewClass(data NewClassRequest) (interface{}, error) {
	var class = Class{
		ClassName: data.ClassName,
		SubjectId: data.SubjectId,
		UserId:    data.UserId,
		CreatedAt: time.Now(),
	}
	result := config.DB.Create(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return class, nil
}

func UpdateClass(id uint, data UpdateClassRequest) (interface{}, error) {
	var class = Class{
		ClassName: data.ClassName,
		SubjectId: data.SubjectId,
		UserId:    data.UserId,
	}
	result := config.DB.Where("id = ?", id).Updates(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return class, nil
}
