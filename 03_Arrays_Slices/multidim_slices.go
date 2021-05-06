package main

import "fmt"

// Like arrays, slices are always 1-dimensional
// but may be composed to construct higher-dimensional objects
// With slices of slices, the length may be dynamic
// Hence the multi dim array can be jagged
// Individual slices must be allocated individually

func main() {

	matrix := make([][]int, 0)

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	matrix = append(matrix, row1)
	matrix = append(matrix, row2)

	fmt.Println("Matrix")
	fmt.Println(matrix)
}
