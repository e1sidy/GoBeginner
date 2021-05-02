package main

import (
	"fmt"
)

var number int = 1

func foo() {

	fmt.Printf("global variable number inside foo function is %d\n", number)
	// local number shadows the gobal scoped number
	var number = 5
	fmt.Printf("local variable number inside foo function is %d\n", number)
}

func main() {

	number = 10
	fmt.Printf("global scoped variable number inside main function is %d\n", number)
	foo()
	fmt.Printf("global scoped variable number inside main function is %d\n", number)
}
