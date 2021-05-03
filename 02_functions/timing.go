package main

import (
	"fmt"
	"time"
)

func calculation() {
	for i := 0; i < 100000; i += 1 {
		//do something
	}
}

func main() {
	start := time.Now()
	calculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("The executing took time - ", delta)
}
