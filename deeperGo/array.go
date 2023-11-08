package main

import "fmt"

func Array() {

	fmt.Println("Array start from here")
	arrOfInt := [6]int{21, 3, 5, 23, 5,4}
	fmt.Printf("%v and %T \n", arrOfInt, arrOfInt)

	arrofString := [4]string{"Ridesh", "nath", "ji"}

	fmt.Printf("%v and %T\n", arrofString, arrofString)

	arrofFloat := [4]float32{3.55, 34.533, 55.33}

	fmt.Printf("%v and %T\n", arrofFloat, arrofFloat)

	fmt.Println("Apply loop on array ")
	max := 0
	min := arrOfInt[0]
	for i := 0; i < len(arrOfInt); i++ {
		if max < arrOfInt[i] {
			max = arrOfInt[i]
		}
		if min > arrOfInt[i] {
			min = arrOfInt[i]

		}
	}

	fmt.Println("maximum value of the array is", max)
	fmt.Println("minimum value of the array is", min)

}
