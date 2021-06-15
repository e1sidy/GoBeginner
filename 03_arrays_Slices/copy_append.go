package main

import "fmt"

func printSlc(slice []int) {
	for _, val := range slice {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

func loadSlc(slice []int, fact int) {
	for i := range slice {
		slice[i] = fact * i
	}
}

func main() {

	slc_to := make([]int, 10)
	slc_from := []int{1, 2, 3}

	// Loading values till index 0 to 3 into slc_to
	loadSlc(slc_to[:4], 5)
	printSlc(slc_to)

	// Copying value of slc_from to slc_to
	copy(slc_to, slc_from)

	printSlc(slc_to)

	// Appending multiple to slc_from
	slc_from = append(slc_from, 4, 5, 6)
	printSlc(slc_from)
	slc1 := []int{7, 8, 9}
	// ... operator used to expand the slice
	slc_from = append(slc_from, slc1...)
	printSlc(slc_from)

}
