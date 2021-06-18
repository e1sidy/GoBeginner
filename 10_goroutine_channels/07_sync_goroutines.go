package main

/*
Closing a channel
- Closing a channel is not necessary
- Only the sender should close a channel, never the receiver.
- It is good practice to do this with a defer statement immediately after the making of the channel

Checking if a channel is closed
	if v, ok := <-ch; ok { ... }
*/

// Modifing the program from 03_channels.go for current shinanigans
import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Guttesburg"
	ch <- "Red Rock"
	ch <- "Delaware"
	ch <- "Pitsburg"
	close(ch)
}

func getData(ch chan string) {
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}
