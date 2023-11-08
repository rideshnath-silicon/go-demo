package main

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Profession string
	Doj        string
	salary     float32
}

func FunctionReciever() {
	Emp1 := Employee{"Ridesh", 22, "Developer", "01/11/2023", 2.48}
	Emp2 := Employee{"Dev", 24, "Developer", "01/11/2023", 2.48}

	fmt.Println(Emp1)
	fmt.Println(Emp2)

	Emp1.upEmpDetails("Ridesh Ji", 23, "Developer", "01/11/2023", 2.48)

	fmt.Println("Employee after update the data ", Emp1)

	print(Emp1)

}

func (q *Employee) upEmpDetails(Name string, age int, prof string, doj string, salary float32) {
	q.Name = Name
	q.Age = age
	q.Profession = prof
	q.salary = salary
	q.Doj = doj
}

func (obj Employee) Greetings() string {
	return obj.Name
}

func print(s GetOutput) {
	greet := s.Greetings()

	fmt.Println(greet)

}

type GetOutput interface {
	Greetings() string
}
