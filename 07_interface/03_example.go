package main

// Program to demonstrate power of interface in handling data of unknown type
// When dealing with data of an unknown type from external sources,
// type testing and conversion to Go data types can be very useful,
//  e.g. parsing data that are JSON- or XML-encoded.

import (
	"fmt"
)

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d's type is unknown\n", i)
		}
	}
}

func main() {
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}
