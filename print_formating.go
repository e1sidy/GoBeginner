package main

import (
	"fmt"
	"math/rand"
)

// %d specifies format for integral values.
// %s specifies format for string values.
// %v specifies the general default format.
// %c specifies ASCII character format.
// %b specifies format for bit representation of number

func main() {
	var c1 complex64 = 5 + 2i
	a := rand.Int()
	b := rand.Intn(10)

	fmt.Printf("complex number %v\n", c1)
	fmt.Printf("Two random numbers are %d and %d\n", a, b)
	fmt.Printf("Bit representation of b is %b\n", b)

	var char1 byte = 'A'
	var char2 byte = 65

	fmt.Printf("ch1 and ch2 are equal %c and %c\n", char1, char2)

	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch1, ch2, ch3)   // UTF-8 code point
}
