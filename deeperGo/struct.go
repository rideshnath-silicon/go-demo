package main

import "fmt"

type Person struct {
	name       string
	age        int
	profession string
}

func structs() {
	fmt.Println("Struct is start")
	// struct is use to store multiple value or diffirent data type  in one variable
	// we can use one struct multiple time to store data

	person1 := Person{"Ridesh", 22, "Developer"}

	fmt.Println(person1.name)
	fmt.Println(person1)
	person2 := Person{"Dev", 24, "Developer"}

	fmt.Println(person2.name)
	fmt.Println(person2)

	var pers1 Person

	pers1.name = "Hege"
	pers1.age = 45

	fmt.Println(pers1)
	pers1.updateAge(55)

	fmt.Println(pers1)
}

func (pers1 *Person) updateAge(upage int) {

	pers1.age = upage

}
