package main

import "fmt"

/** Slices is reference type memory structure with 3 fields:
1- a pointer to underlying array
2- the lenght of slice	len(s)
3- capacity of slice	cap(s)

Capacity of the slice => lenght of slice + the length of array beyond the slice
0 <= len(s) <= cap(s)

Since arrays in GO are not like C, the pointer to arrays do not point to only the first element

Multiple slices can point to pieces of same array but multiple array can never share data
As slices are references, they don’t use up additional memory
and they are more efficient to use than arrays.
*/

func foo(slice []int) {
	fmt.Println("Slice ->", slice)
	fmt.Println("Length ->", len(slice))
	fmt.Println("Capacity ->", cap(slice))
}

func main() {

	var arr [6]int
	var slice = arr[2:5]

	for i := range arr {
		arr[i] = i
	}

	fmt.Printf("The size of the array is %d\n", len(arr))
	foo(slice)

	// Growing the slice
	slice = slice[0:4]
	foo(slice)
}
