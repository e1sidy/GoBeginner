package main

import "fmt"

func main() {
	fv := func() {
		fmt.Println("Hello Closure!")
	}
	fv()
	fmt.Printf("Type of function literal/closure : %T\n", fv)
}
