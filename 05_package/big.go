package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt32)
	io := big.NewInt(1928)
	io.Mul(io, im)
	fmt.Printf("The value of multiplying is: %v\n", io)

	ir := big.NewRat(984, math.MaxInt32)
	ir2 := big.NewRat(math.MinInt32, 342)
	ir3 := big.NewRat(1, 1)
	ir3.Add(ir, ir2)
	fmt.Printf("The value for addition is: %v\n", ir3)
}
