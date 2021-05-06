package main

import "fmt"

func main() {
	var arr [5]int
	for i := range arr {
		arr[i] = 2 * i
		fmt.Printf("The value of element at %d index is %d\n", i, arr[i])
	}
}
