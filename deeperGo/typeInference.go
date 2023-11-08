package main

import "fmt"

func typeInference() {
	fmt.Println("Type interface start")

	// type interface use to declare multple variable together with having different datatype

	var name, age, profession, isTechnincal, salary = "Ridesh", 22, "Developer", true, 2.48

	fmt.Printf("The name has value %v and type %T\n", name, name)
	fmt.Printf("The age has value %v and type %T\n", age, age)
	fmt.Printf("The profession has value %v and type %T\n", profession, profession)
	fmt.Printf("The isTechnincal has value %v and type %T\n", isTechnincal, isTechnincal)
	fmt.Printf("The salary has value %v and type %T\n", salary, salary)
}
