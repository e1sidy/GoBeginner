package main

import (
	"flag"
	"fmt"
)

func iter(b int, ch chan int) {
	for i := 0; i < b; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	n := flag.Int("num", 42, "a int")
	flag.Parse()
	a := make(chan int)
	b := make(chan int)

	go iter(*n, a)
	go iter(*n, b)

	for a != nil || b != nil {
		select {
		case x, ok := <-a:
			if ok {
				fmt.Println(x)
			} else {
				a = nil
			}
		case y, ok := <-b:
			if ok {
				fmt.Println(y)
			} else {
				b = nil
			}
		}
	}
}
