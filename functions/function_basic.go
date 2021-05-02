package main

import "fmt"

func f2(a, b, c int) int {
	return a + b + c
}

func f1(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func main() {
	fmt.Printf("The result is %d\n", f2(f1(10, 20)))
}
