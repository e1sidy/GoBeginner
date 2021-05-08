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

	// Consider the following example
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	// The order of execution is from left to right
	// 1- rq = (rm*rn)
	// 2- rq = rq+ro
	// 3- rq = rq*rp
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}
