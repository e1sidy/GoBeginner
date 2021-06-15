package main

import "fmt"

func printArr(arr [5]int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}

func multBy2(arr [5]int) {
	for i := range arr {
		arr[i] = 2 * arr[i]
	}
}

func printArrPtr(arr *[5]int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}

func multBy2Ptr(arr *[5]int) {
	for i := range arr {
		arr[i] = 2 * i
	}
}

func main() {

	// using new to create a array return pointer to that array
	// var arrPtr = new([5]int)
	var arr [5]int

	printArr(arr)
	multBy2(arr)
	printArr(arr)

	printArrPtr(&arr)
	multBy2Ptr(&arr)
	printArrPtr(&arr)

}
