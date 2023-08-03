package main

import (
	"fmt"
)

func arrr() {
	//primitive arrays
	students := [...]int{43, 44, 53}
	fmt.Printf("%v, %T\n", students, students)
	var people [3]string
	people[0] = "Jonathan"
	people[1] = "Mariam"
	people[2] = "Nathan"
	fmt.Printf("%v,%T\n", people, people)
	fmt.Printf("%v,%T\n", people[0], people[1])
	var streets [3]string = [3]string{"Meri", "Summit", "CMC"}
	fmt.Printf("%v,%T\n", streets, streets)
	//Non primitive Arrays
	var matrix [3][3]int
	matrix[0] = [3]int{0, 1, 2}
	matrix[1] = [3]int{3, 4, 5}
	matrix[2] = [3]int{6, 7, 8}
	fmt.Printf("%v\n", matrix)

	originalArray := [3]int{1, 2, 3}
	copyArray := originalArray
	copyArray[0] = 5
	fmt.Println(originalArray, copyArray) // changing the copy doesn't change the original, because the copy doesn't point to it
	pointArray := &originalArray
	pointArray[0] = 25
	fmt.Println(originalArray, pointArray) //changig the point array affects the origial array

}
