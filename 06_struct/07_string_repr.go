package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func (this *TwoInts) String() string {
	return "(" + strconv.Itoa(this.a) + " / " + strconv.Itoa(this.b) + ")"
}

func main() {
	obj := new(TwoInts)
	obj.a = 29
	obj.b = 90

	fmt.Println(obj)
}
