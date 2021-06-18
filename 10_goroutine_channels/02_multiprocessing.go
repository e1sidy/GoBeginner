package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	numCores := flag.Int("n", 2, "number of CPU cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
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
