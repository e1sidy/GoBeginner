package main

import "fmt"

type Stringer interface{ String() string }

type Class struct {
	val string
}

func (this *Class) String() string {
	return this.val
}

func main() {

	v := Stringer(&Class{"add"})

	// Check if a type implements Stringer interface
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}
