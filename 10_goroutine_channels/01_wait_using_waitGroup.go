package main

import (
	"fmt"
	"sync"
)

func HeavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 1")
	// Do a lot of stuff
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 2")
	// Do a lot of stuff
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go HeavyFunction1(wg)
	go HeavyFunction2(wg)
	wg.Wait()
	fmt.Println("All Finished!")
}
