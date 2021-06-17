package main

import (
	"flag"
	"fmt"
)

// Ways to run using flags
// go run main.go
// Type go run main.go -h
// Type go run main.go -lang="V" -num=42

func main() {
	strPtr := flag.String("lang", "Go", "a string")
	numPtr := flag.Int("num", 108, "an int")
	boolPtr := flag.Bool("truth", false, "a bool")
	var str string
	flag.StringVar(&str, "str", "Crystal", "a string variable")
	// Parse all the flags defined, must be called before using any flag
	flag.Parse()

	fmt.Println("lang:", *strPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("truth:", *boolPtr)
	fmt.Println("str:", str)
	fmt.Println("tail:", flag.Args())
}
