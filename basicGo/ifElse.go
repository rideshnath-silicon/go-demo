package main

import "fmt"

func conditional() {
	fmt.Print("\n Now start Conditional program from here :- \n")

	a := 10
	b := 20
	c := 80

	output := checkGreater(a, b, c)

	fmt.Println(output)
}



func checkGreater(a int, b int, c int) int {
	if a > b {
		if a > c {
			fmt.Println(" A is greater than b and c")
		} else {
			fmt.Println(" c is greater than a and b")
		}

	} else if b > a {
		if b > c {
			fmt.Println(" b is greater than a and c")
		} else {
			fmt.Println(" c is greater than a and b sec")
		}
	} else {
		fmt.Println(" c is greater than a and b")
	}

	sumOfAllArg := sum(a, b, c)
	return sumOfAllArg
}


func sum(a int, b int, c int) int {

	fmt.Print("sum of All ARG is :- ")
	return a + b + c
}