package main

import "fmt"

// Two channels one receiving input channel and one output channel for sending

// Find prime numbers using pipe and filter pattern
// implementing sieve, generate and filter as factories
// factories are function that make a channel and return sit and they use lambda functions as goroutines

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func seive() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := seive()
	for {
		fmt.Println(<-primes)
	}
}
