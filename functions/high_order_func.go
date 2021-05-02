package main

import (
	"fmt"
)

func isOdd(num int) bool {
	return num%2 == 1
}

func isEven(num int) bool {
	return num%2 == 0
}

func filterSlice(sl []int, flt func(int) bool) []int {
	var res []int
	for _, val := range sl {
		if flt(val) {
			res = append(res, val)
		}
	}

	return res

}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filterSlice(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filterSlice(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
