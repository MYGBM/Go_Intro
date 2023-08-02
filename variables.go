package main

import (
	"fmt"
	"strconv"
)

func convert() {
	i := 5
	fmt.Printf("%v,%T\n", i, i)
	var j float32 = 55.5
	fmt.Printf("%v,%T\n", j, j)
	var k int
	k = int(j)
	fmt.Printf("%v,%T\n", k, k)
	var l float64
	l = float64(i)
	fmt.Printf("%v,%T\n", l, l)
	var name string
	name = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", name, name)

}
