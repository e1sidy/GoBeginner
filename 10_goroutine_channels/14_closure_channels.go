package main

// Since functions are values in go they can be send through a closure

import (
	"fmt"
)

func prod(ch chan func()) {
	f := <-ch
	f()
	ch <- nil
}

func main() {
	ch := make(chan func())
	go prod(ch)
	x := 0
	ch <- func() {
		x++
	}
	fmt.Println(x)
}
