package main

import (
	"encoding/json"
)

func marshaling(id int) string {
	var employee Employee

	result := db.Where("ID = ?",id).Find(&employee)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	JsonData, Err := json.Marshal(employee)

	if Err != nil {
		panic("Failed to encode in json" + Err.Error())
	}
	
	return string(JsonData)
}
