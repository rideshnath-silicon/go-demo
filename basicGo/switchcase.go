package main

import "fmt"

func switchcase() {
	fmt.Print("\n Now start switch case from here :- \n")
	var swit int
	fmt.Print("enter int value for switch case  ")
	fmt.Scanf("%d", &swit)
	println(swit)
	switch swit {
	case 1, 6:
		fmt.Println("First case ")
	case 2:
		fmt.Println("second case ")
	case 3:
		fmt.Println("third case ")
	case 4:
		fmt.Println("fourth case ")
	default:
		fmt.Println("No case Found")
	}
	
}
