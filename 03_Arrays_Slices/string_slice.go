package main

import "fmt"

func main() {
	s := "Aabc"
	s1 := "9a0fdj"

	// Slice of a string can be made by converting the string to byte
	slc := []byte(s)
	slc = append(slc, 'a')

	// Appending a string to byte slice
	slc = append(slc, s1...)
	fmt.Println(slc)

	// Making sub-string
	subStr := slc[1:5]
	fmt.Println(subStr)

	// Slices can be used to mutate strings
	s2 := "Hello World!"
	slc2 := []byte(s2)
	slc2[len(slc2)-1] = 's'
	s2 = string(slc2)
	fmt.Println(s2)

}
