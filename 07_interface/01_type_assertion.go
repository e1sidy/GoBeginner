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

func (this Circle) Area() float32 {
	return math.Pi * this.radius * this.radius
}

func main() {
	sq := &Square{8}
	areaIntf := Shaper(sq)

	// One can see that since *Square implements Area so we can call interface on *Square
	// in the case of Circle, we can call interface on Circle only, but not *Circle
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is %T\n", t)
	}
	if t, ok := areaIntf.(Circle); ok {
		fmt.Printf("The type of areaIntf is %T\n", t)
	} else {
		fmt.Println("areaIntf is not of type implementing interface")
	}

}
