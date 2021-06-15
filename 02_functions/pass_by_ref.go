package main

import (
	"fmt"
)

func factorial(n int, fact *int) {
	if n == 0 {
		return
	}
	*fact *= n
	factorial(n-1, fact)
}

func main() {
	var n int = 5
	var fact int = 1

	factorial(n, &fact)

	fmt.Printf("The factorial of n is :%d\n", fact)

}
