package main

// It is a producer that only returns the next value, not the entire sequence. This is called _lazy evaluations, which means only computing what you need at the moment, therefore saving valuable resources

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0

	go func() {
		for {
			yield <- count
			count++
		}
	}()

	return yield
}

func generateIntegers() int {
	return <-resume
}

func main() {
	resume = integers()
	fmt.Println(generateIntegers())
	fmt.Println(generateIntegers())
	fmt.Println(generateIntegers())

}
