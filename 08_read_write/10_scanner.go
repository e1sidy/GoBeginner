package main

import (
	"bufio"
	"fmt"
	"os"
)

func example1() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter some input:")
	// scanner.Split sets the default split funciton for scanner
	// default split function is bufio.ScanLines
	scanner.Split(bufio.ScanWords)

	// Scan() advances to the next token which can be accessed
	// using Text() or Bytes()
	for scanner.Scan() {
		fmt.Printf("Token ->%s<-\n", scanner.Text())
	}
}

func main() {
	example1()
}
