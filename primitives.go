package main

import (
	"fmt"
)

func primitives() {
	//integers
	var unsigned_Number uint = 45
	fmt.Printf("%v,%T\n", unsigned_Number, unsigned_Number)
	var byteInteger int8 = 10
	fmt.Printf("%v,%T\n", byteInteger, byteInteger)
	var n int
	fmt.Printf("%v,%T\n", n, n)
	//bitwise operations
	a := 10
	b := 3
	fmt.Println(a & b)  //both
	fmt.Println(a | b)  // either or both
	fmt.Println(a ^ b)  // one and only one
	fmt.Println(a & ^b) // neither
	//bit shifting
	num1 := 8
	num2 := 16
	fmt.Println(num1 >> 3) //2^3/2^3
	fmt.Println(num2 << 3) //2^3*2^3

	//floats
	float1 := 32.8
	float1 = 10.2e2
	float1 = 13.5e7
	fmt.Printf("%v,%T\n", float1, float1)

	//ComplexNumbers
	var complexNum complex64 = 2i
	fmt.Printf("%v,%T\n", complexNum, complexNum)

	//Text Types

	var text string = "banana"
	fmt.Printf("%v,%T\n", text, text)
	fmt.Printf("%v,%T\n", text[0], text[0])
	fmt.Printf("%v,%T\n", string(text[0]), text[0])
	byteSlicer := []byte(text)
	fmt.Printf("%v,%T\n", byteSlicer, byteSlicer)

}
