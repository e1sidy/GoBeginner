package main

import "fmt"

func sliceDetails(slice []int) {
	fmt.Println("Slice ->", slice)
	fmt.Println("Length ->", len(slice))
	fmt.Println("Capacity ->", cap(slice))
}

func foo(slice []int, fact int) {
	for i := range slice {
		slice[i] = fact * i
	}
}

func main() {

	// Making a slice with make
	var slc1 []int = make([]int, 4)

	foo(slc1, 3)
	sliceDetails(slc1)

	// using make with ([]T, len, cap)
	slc2 := make([]int, 5, 10)
	foo(slc2, 8)
	sliceDetails(slc2)

	// Difference between make() and new()
	// Though both allocate memory in heap they do different things and apply to different types
	// new(T) ->
	//	1.	It allocates zeroed storage for a new item of type T
	// 		and returns its address as a value of type *T
	//  2. 	It applies to value types like arrays and structs
	//
	// make(T) ->
	//	1.	It returns an initialized value of type T
	//	2.	It applies to three reference types - slices, maps and channels

	// They are equivalent in the slice is of len(s) = 0
	sliceDetails(*new([]int))
	sliceDetails(make([]int, 0))
}
