package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	if input, err := inputReader.ReadString('\n'); err == nil {
		fmt.Printf("The input was %s\n", input)
	}
}
