package main

import "fmt"

// Maps in GO are a reference type data structure
// Declaration -->  var map1 map[keytype]valuetype
// keytype can be any type for which == and != is defined
// 	this include int, float, string
// 	In Go for array and struct equality operator is defined by element-wise comparision
// 	So array, struct, pointer and interface can also be used as keyvalue
//  But slices can not be used as equality is not defined for them

// Map can be used to efficently pass value to a function as they are reference type

// Maps are reference types, as memory is allocated with the make function

// Note: Don’t use new, always use make with maps.

func main() {

	var mapLit map[string]int = map[string]int{"Ine": 1, "Dui": 2, "Trie": 3}
	var mapCreated map[string]float32 = make(map[string]float32)
	var mapAssigned map[string]int = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Println(mapAssigned)
	fmt.Println(mapCreated)

	// Map can take any type of value
	// Example:
	mf := map[int]func() int{ // key type int, and value type func()int
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	// Using slices as value for maps
	// Example:
	// mp1 := make(map[int][]int)

}
