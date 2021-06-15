package main

import (
	"fmt"
)

func main() {
	val := 42
	if val == 42 {
		fmt.Println("The secret of Universe")
	} else {
		fmt.Println("It is a normal number")
	}

	if num := 42; num > 50 {
		fmt.Println("True block")
	} else {
		fmt.Println("False block")
	}
}
