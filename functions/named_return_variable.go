package main

import "fmt"

func namedReturn(a int, b int) (sum, diff, prod int, div float32) {
	sum, diff, prod, div = a+b, a-b, a*b, float32(a)/float32(b)
	return
}

func main() {

	val1, val2, val3, val4 := namedReturn(30, 20)

	fmt.Printf("The return value are %d, %d, %d, %.2f\n", val1, val2, val3, val4)
}
