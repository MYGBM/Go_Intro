package main

import (
	"fmt"
)

const ExportedConst = "Exported" // always exportable same package in d/f packages UpperCase=Exportable, LowerCase=Unexportable

func coonstants() {
	//Typed Constants must be of same type for operations
	// const a int8 = 42
	// const b int16 = 24
	// 	fmt.Printf("%v,%T\n", a+b, a+b) error

	//Visibility
	const notVisible string = "NotVisible"
	fmt.Printf("%v,Visible inside Function", notVisible)

	const c = 42 // Untyped
	const d int32 = 24
	var e int8 = 4
	fmt.Printf("%v,%T\n", c+d, c+d) // can interoperate w/h d/df types
	fmt.Printf("%v,%T\n", c+e, c+e)

	//enumerated constants
	const (
		Red = iota
		Green
		Blue
	)
	fmt.Printf("%v,%T\n", Red, Red)
	fmt.Printf("%v,%T\n", Green, Green)
	fmt.Printf("%v,%T\n", Blue, Blue)

	// Enumerated Expressions
	const (
		Read    = 1 << iota //1
		Write               // 2
		Execute             // 4

	)
	var perm int = Read | Write
	fmt.Println(perm)

}
