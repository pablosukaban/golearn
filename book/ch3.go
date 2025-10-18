package main

import (
	"fmt"
	"maps"
	"slices"
)

func printIntArrCap(arr []int) {
	fmt.Println("arr: ", arr, len(arr), cap(arr))
}
func printStrArrCap(arr []string) {
	fmt.Println("arr: ", arr, len(arr), cap(arr))
}

func SlicesTest() {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	z := []int{1, 2, 3, 4}

	var n []int
	// w := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	fmt.Println(n == nil)
	// fmt.Println(slices.Equal(x, w))

	x = append(x, 4)
	fmt.Println(x)
}

func CapacityTest() {
	var x []int
	printIntArrCap(x)
	x = append(x, 10)
	printIntArrCap(x)
	x = append(x, 20)
	printIntArrCap(x)
	x = append(x, 30)
	printIntArrCap(x)
	x = append(x, 40)
	printIntArrCap(x)
	x = append(x, 50)
	printIntArrCap(x)
}

func MakeTest() {
	x := make([]int, 4)
	x = append(x, 10)
	printIntArrCap(x)

	y := make([]int, 0)
	y = append(y, 4, 5, 6)
	printIntArrCap(y)

	z := []int{1, 2, 3, 4}
	clear(z)
	printIntArrCap(z)
}

func SlicingSliceTest() {
	x := []string{"a", "b", "c", "d"}
	printStrArrCap(x[:2])
	printStrArrCap(x)
	y := x[:2]
	printStrArrCap(y)
	y = append(y, "z")
	fmt.Println("appended")
	printStrArrCap(x)
	printStrArrCap(y)
}

func MapsTest() {
	x1 := map[string]int{
		"yo": 1,
	}
	x2 := map[string]int{
		"yo": 1,
	}

	fmt.Println(maps.Equal(x1, x2))
}

func MapsAsSetTest() {
	intSet := map[int]bool{}

	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, val := range vals {
		intSet[val] = true
	}
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[12])

	// ================

	intSetStructs := map[int]struct{}{}

	for _, val := range vals {
		intSetStructs[val] = struct{}{}
	}

	if _, ok := intSetStructs[5]; ok {
		fmt.Println("5 in there")
	}

	if _, ok := intSetStructs[1000]; !ok {
		fmt.Println("1000 not in there")
	}
}

func StructTest() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println(fred)
	tom := person{}
	fmt.Println(tom)
	julia := person{
		"Julia",
		40,
		"finik",
	}
	fmt.Println(julia.pet)
	beth := person{
		name: "Beth",
		pet:  "petName",
	}
	fmt.Println(beth)

	var anonPerson struct {
		name string
		age  int
		pet  string
	}
	anonPerson.name = "anon"
	anonPerson.age = 30
	fmt.Println("anonPerson", anonPerson)

	anonPet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println("anonPet", anonPet)
}
