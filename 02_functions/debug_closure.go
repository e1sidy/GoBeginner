package main

import (
	"fmt"
	"log"
	"runtime"
)

func runtimeCls() {
	where := func() { // debugging function
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	a := 2 * 5
	where()
	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}

func logCls() {
	var where = log.Print

	where()
	// some code
	a := 2 * 5
	where()

	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}

func main() {
	fmt.Println("Debug with runtime and closure")
	runtimeCls()
	fmt.Println("Debug with log and closure")
	logCls()
}
