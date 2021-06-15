package main

import (
	"fmt"
)

func main() {
	var num1 int = 99

	// Switch Case example
	switch num1 {
	case 99, 100:
		fmt.Println("The number is 99 or 100")
	case 98:
		fmt.Println("The number is 98")
	default:
		fmt.Println("Default Case")
	}

	// Switch statement with condition
	// switch statement without a variable
	// is known as a switch true
	switch {
	case num1 < 0:
		fmt.Println("The number is negative")
	case num1 < 100:
		fmt.Println("The number is less than 100")
	default:
		fmt.Println("The number is greater than 100")
	}

	// Switch with initialization
	var t int
	switch a, b := 10, 50; {
	case a < b:
		t = -1
	case a == b:
		t = 0
	case a > b:
		t = 1
	}
	fmt.Println("The value of t is ", t)
}
