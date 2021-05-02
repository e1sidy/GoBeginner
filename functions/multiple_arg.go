package main

import (
	"fmt"
)

func expandArg(a ...int) {
	if len(a) == 0 {
		return
	}

	for _, val := range a {
		fmt.Printf("%d \t", val)
	}

	fmt.Println()
}

func main() {
	expandArg(1, 3, 5, 7, 9, 13)
}
