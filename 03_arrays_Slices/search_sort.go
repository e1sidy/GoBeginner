package main

import (
	"fmt"
	"sort"
)

func main() {
	slcInt := []int{24, 314, 14, 1, 135}
	slcFloat := []float64{1134.1, 13.43, 135.14, 134.14}

	sort.Ints(slcInt)
	sort.Float64s(slcFloat)
	fmt.Println(slcInt)
	fmt.Println(slcFloat)

	// Returns the index if found else
	// return value is the index to insert x if x is not present
	fmt.Println(sort.SearchInts(slcInt, 30992))
	fmt.Println(sort.SearchInts(slcInt, 314))
}
