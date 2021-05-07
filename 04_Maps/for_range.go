package main

import (
	"fmt"
)

func main() {
	m1 := make(map[int]float32, 10)
	m1[1] = 1.0
	m1[2] = 2.0
	m1[3] = 3.0
	m1[4] = 4.0

	// for range construct used to get the keys
	for key := range m1 {
		fmt.Println(key)
	}

	// for range construct used to get the key value pairs
	for key, value := range m1 {
		fmt.Println(key, "->", value)
	}
}
