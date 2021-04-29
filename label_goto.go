package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
LOOP:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break LOOP
		}
		fmt.Println(i)
	}
}
