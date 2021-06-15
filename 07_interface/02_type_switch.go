package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (this *Square) Area() float32 {
	return this.side * this.side
}

func (this *Circle) Area() float32 {
	return math.Pi * this.radius * this.radius
}

func main() {
	sq := &Square{8}
	areaIntf := Shaper(sq)

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)

	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)

	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
