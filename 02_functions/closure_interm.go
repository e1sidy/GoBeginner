package main

import (
	"fmt"
)

type (
	arg1r1Cls func(int) int
)

// The key idea begin closure is that the return function
// has access to the local variable of the enclosing function

// If the local variable changes the closure function will have
// access to the updated value of the local variable
// See cummlator function below for example

func main() {
	add23 := adder(23)
	fmt.Printf("add23(56) = %d, \n", add23(56))
	fmt.Printf("add23(117) = %d, \n", add23(117))

	// Calling cummlator function
	// As in the output the variable x is accessable to f
	// and is increase by delta after each call
	// the change is reflected on each subsequent call
	f := cummlator()
	fmt.Print(f(1), ", ")
	fmt.Print(f(20), ", ")
	fmt.Print(f(300), "\n")

	// Fibonacci using closure
	fib := fibCls()
	fmt.Println("Printing fibonacci sequence")
	for i := 0; i < 10; i += 1 {
		fmt.Print(fib(), ", ")
	}

	// Filter Factory function
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("\ns = ", s)
	odd_even_function := filter_factory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)
	//separate those that are bigger than 4 and those that are not.
	bigger, smaller := filter_factory(isGreaterThan4)(s)
	fmt.Println("Bigger than 4: ", bigger)
	fmt.Println("Smaller than or equal to 4: ", smaller)
}

// Normal adder closure usage

func adder(x int) arg1r1Cls {
	return func(val int) int {
		return x + val
	}
}

// cummlator function to show that closures can be used to store states as a class

func cummlator() arg1r1Cls {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// fibonacci using closure
func fibCls() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func isOdd(x int) bool {
	return x%2 == 1
}

func isGreaterThan4(x int) bool {
	return x > 4
}

// Factory function to split a array into two array based on predicate

func filter_factory(f func(int) bool) func([]int) ([]int, []int) {
	return func(arr []int) (tArr []int, fArr []int) {
		for _, val := range arr {
			if f(val) {
				tArr = append(tArr, val)
			} else {
				fArr = append(fArr, val)
			}
		}
		return
	}
}
