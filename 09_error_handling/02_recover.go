package main

// Recover, built-in-function, allows a program to regain control of a panicking Go routine
// The recover is only useful when called inside a deferred function
// It then retrieves the error value passed through the call to panic
// When used in normal execution, a call to recover will return nil and have no other effect
// The panic causes the stack to unwind until a deferred recover() is found or the program terminates.
// It is analogous to the catch block in the Java and .NET languages.

/*
Here are a couple of best practices which every writer of custom packages should apply:

- Always recover from panic in your package: no explicit panic() should be allowed to cross a package boundary.
- Return errors as error values to the callers of your package.
*/
import (
	"fmt"
)

func badCall() {
	a, b := 0, 10
	n := b / a
	fmt.Println("No panic, result", n)
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicking %s\n", err)
		}
	}()
	badCall()
	fmt.Println("This should not execute")
}

func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test Completed")
}
