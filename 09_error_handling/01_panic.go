package main

// Go Panicing
// - If panic is called from a nested function, it immediately stops the execution of the current function
// - All defer statements are guaranteed to execute.
// - This bubbles up to the top level, executing defers, and at the top of the stack, the program crashes.
// - Temination of program from panicing should be avoided, and recovery should be attempted whenever possible

import (
	"fmt"
	"os"
)

func foo() {
	fmt.Println("Defer function always perform the cleanup even after panicing starts")
}

func main() {
	fmt.Println("Starting the Program")
	defer foo()
	if len(os.Args) > 1 {
		panic("A severe error occurred: stopping the program!")
	}
	fmt.Println("Ending the program")
}
