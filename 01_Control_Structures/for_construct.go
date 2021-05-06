package main

import (
	"fmt"
)

func main() {
	str := "Go is a different language"
	fmt.Println("The size of str is ", len(str))

	for idx := 0; idx < len(str); idx += 1 {
		fmt.Printf("The character at %d location is %c\n", idx, str[idx])
	}

	// This causes error if we use non-ASCII characters
	// There character size is 2 to 4 byte
	str2 := "Chinese: 日本語"
	fmt.Println("The size of str is ", len(str2))

	for idx := 0; idx < len(str2); idx += 1 {
		fmt.Printf("The character at %d location is %c\n", idx, str2[idx])
	}

}
