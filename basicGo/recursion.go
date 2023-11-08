package main

import "fmt"

func recursion(){
	fmt.Println("Exmaple countdown for recursive function")
	countDown(10)
}

func countDown(a int){
	if a > 0 {
		fmt.Println("Value of a is :- ",a)
		countDown(a-1)
	}else {
		fmt.Println("countdown Stop ")
	}

}