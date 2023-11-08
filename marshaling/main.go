package main

import "fmt"

var db = ConnectDatabase()

func main() {
	EmployeeData := marshaling(1)

	// fmt.Println(EmployeeData)

	var empl Employee

	data := unmarshal(EmployeeData, empl)

	fmt.Println(data)
}
