package main

import "fmt"

// To create map of slice use make in two steps
// First, to create the outer slice
// Second, to create individual map in slice

func main() {
	maps := make([]map[int]int, 5)
	for idx := range maps {
		maps[idx] = make(map[int]int, 1)
		maps[idx][1] = 5 * idx
	}

	for _, m := range maps {
		fmt.Println(m)
	}
}
