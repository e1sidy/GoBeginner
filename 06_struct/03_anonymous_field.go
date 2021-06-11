package main

import (
	"fmt"
)

type innerS struct {
	i1 int
	i2 int
}

type outerS struct {
	i1 int
	f1 float32
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.i1 = 5
	outer.f1 = 8.3
	// if there is a conflict then refer to the inner struct to resolve conflict
	outer.innerS.i1 = 7
	outer.i2 = 923
	outer.int = 0

	fmt.Println(outer)

}
