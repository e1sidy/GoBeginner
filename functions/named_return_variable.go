package main

import "fmt"

func namedReturn(a int, b int) (sum, diff, prod int, div float32) {
	sum = a + b
	diff = a - b
	prod = a * b
	div = float32(a) / float32(b)
	return
}

func main() {

	val1, val2, val3, val4 := namedReturn(20, 10)

	fmt.Printf("The return value are %d, %d, %d, %.2f\n", val1, val2, val3, val4)
}
