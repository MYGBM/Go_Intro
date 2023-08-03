package main

import (
	"fmt"
)

func slicess() {

	//Slices are Reference Types unlike arrays which are Value Types
	a := []int{1, 2, 3, 4, 5}
	b := a
	b[0] = 100
	fmt.Println(a, b)

	//another way of creating slices
	c := a[1:4]
	fmt.Println(c)

	//third way using makig slices
	slice2 := make([]int, 3, 20)
	fmt.Println(slice2)
	var slice3 = slice2
	slice3[0] = 7
	fmt.Println(slice3, slice2)
	slice2 = append([]int{1, 2, 3})
	fmt.Println(slice2)

}
