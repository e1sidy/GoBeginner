package main

/*
- Data is passed around on channels: only one goroutine has access to a data item at anhy given time
- General Format:
- var identifier chan datatype
- Default value for uninitialized channel is nil
- A channel is, in fact, a typed message queue: data can be transmitted through it.
	It is a First In First Out (FIFO) structure
- A channel is also a reference type, so we have to use the make() function to allocate memory for it.

- Two types of channels - Unbuffered and buffered
- These channels can also be used to implement mutex and semaphore
*/

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go recvData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Guttesburg"
	ch <- "Red Rock"
	ch <- "Delaware"
	ch <- "Pitsburg"
}

func recvData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
	close(ch)
}
