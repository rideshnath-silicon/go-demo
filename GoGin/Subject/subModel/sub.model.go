package submodel

import (
	"gin/config"
	"time"
)

func NewSubject(data AddSubjectRequest) (interface{}, error) {
	var sub = Subject{
		SubName:    data.SubName,
		AuthorName: data.AuthorName,
		Number:     data.Number,
		CreatedAt:  time.Now(),
	}
	result := config.DB.Create(&sub)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return sub, nil
}

func UpdateSubject(Id uint, data UpdateSubjectRequest) (interface{}, error) {
	var sub = Subject{
		SubName:    data.SubName,
		AuthorName: data.AuthorName,
		Number:     data.Number,
	}
	result := config.DB.Updates(&sub)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return sub, nil
}
