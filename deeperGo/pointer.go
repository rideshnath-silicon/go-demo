package main

import "fmt"

func point() {
	a := 55

	var b *int = &a

	fmt.Println(a, *b)

	*b = 54
	fmt.Println(a, *b)

	upValue(&a)

	fmt.Println("After update the value ", a)

}

func upValue(num *int) {
	*num = 100
}
