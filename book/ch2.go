package main

import "fmt"

func Ch2Exercises() {
	// страница 35
	const value = 10

	var i1 float64 = 20
	var f1 float64 = i1

	fmt.Println("1st: ", i1, f1)

	var i2 int = value
	var f2 float64 = value

	fmt.Println("2nd: ", i2, f2)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println("3rd: ", b+1, smallI+1, bigI+1)
}
