package main

import (
	"fmt"
)

var (
	firstName, lastName, str string
	i                        int
	f                        float32
	input                    = "56.12 / 5212 / Go"
	format                   = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("Hi", firstName, lastName)
	fmt.Println("Please enter a new name")
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hello %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &str)
	fmt.Println("From the string we read: ", f, i, str)
}
