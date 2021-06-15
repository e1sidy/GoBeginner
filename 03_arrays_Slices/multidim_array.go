package main

import "fmt"

// Arrays are always 1-dimensional, but they may be composed to form multidimensional arrays,

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int                  // aliasing int as pixel
var screen [WIDTH][HEIGHT]pixel // global 2D array

func main() {
	fmt.Printf("%T\n", screen)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}
