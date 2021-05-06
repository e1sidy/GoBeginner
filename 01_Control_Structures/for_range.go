package main

import (
	"fmt"
)

func main() {
	str1 := "Go is a different language"
	fmt.Println("The size of str is ", len(str1))

	for idx, char := range str1 {
		fmt.Printf("The character at %d location is %c\n", idx, char)
	}

	// This does not causes unexpected behaviour
	// as we are using for range Construct
	str2 := "Chinese: 日本語"
	fmt.Println("The size of str is ", len(str2))

	for idx, char := range str2 {
		fmt.Printf("The character at %d location is %c\n", idx, char)
	}

}
