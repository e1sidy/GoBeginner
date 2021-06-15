package main

import "fmt"

// Passing pointer to array
// as normal pass copies the array and then pass
func printArr(arr *[5]int) {
	// No dereferencing needed for arr(*arr)
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Printf("\n")
}

func main() {

	// Various ways of declaring array literals
	var arr1 = [5]int{3, 6, 7, 3, 7}
	var arr2 = []int{3, 5, 1, 6, 1}
	var arr3 = [...]int{8, 12, 12, 5, 2}
	var arr4 = [5]int{2: 40, 4: 160}

	arr2[1] = 1 // done to remove error of not using

	printArr(&arr1)
	// printArr(&arr2)  // will cause error
	printArr(&arr3)
	printArr(&arr4)
}
