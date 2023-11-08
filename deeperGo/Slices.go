package main

import "fmt"

func Slices() {
	fmt.Println("slices start from here")
	SliceOfInt := []int{21, 3, 5, 23, 5, 4}
	fmt.Printf("%v and %T \n", SliceOfInt, SliceOfInt)
	secSlice := []int{55, 64, 23, 55}
	SliceOfInt = append(SliceOfInt, 3, 5, 3, 2)
	fmt.Println(SliceOfInt)
	fmt.Println(secSlice)
	SliceOfInt = append(SliceOfInt, secSlice...)
	fmt.Println(SliceOfInt)

	max := 0
	min := SliceOfInt[0]
	for i := 0; i < len(SliceOfInt); i++ {
		if max < SliceOfInt[i] {
			max = SliceOfInt[i]
		}
		if min > SliceOfInt[i] {
			min = SliceOfInt[i]

		}
		if SliceOfInt[i] == 5 {
			SliceOfInt[i] = 110

		}
	}

	fmt.Println("maximum value of the slice is", max)
	fmt.Println("minimum value of the slice is", min)
	fmt.Println(len(secSlice))
	fmt.Println((SliceOfInt))

}
