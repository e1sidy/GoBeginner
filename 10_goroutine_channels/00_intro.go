package main

import (
	"fmt"
	"time"
)

// Number of threads used for running a go program is 1 by default
// So even if you use multiple goroutines, it only achives concurency in the code
// Set runtime.GOMAXPROCS = t, to use multiple threads
// Generally number of goroutines > 1 + GOMAXPROCS > 1

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning of longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning of shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait()")
}
