package main

/*
- Using select as 'communication switch' to select values between different channels
- select chooses which of the multiple communications listed by its cases can proceed.
	1- If all are blocked, it waits until one can proceed.
	2- When none of the channel operations can proceed, and the default clause is present, then this is executed because the default is always runnable, or ready to execute.
	3- If multiple can proceed, it chooses one at random.
*/

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received value on channel 1 : ", v)
		case v := <-ch2:
			fmt.Println("Received value of channel 2 : ", v)
		}
	}
}
