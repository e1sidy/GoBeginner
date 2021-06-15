/*
When you call a method on an interface, it must either have an identical receiver type or it must be directly discernible from the concrete type:
1- Pointer methods can be called with pointers.
2- Value methods can be called with values.
3- Value-receiver methods can be called with pointer values because they can be dereferenced first.
4- Pointer-receiver methods cannot be called with values; however, because the value stored inside an interface has no address.
*/

package main

import (
	"fmt"
)

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in function argument:
	// List does not implement Appender (Append method requires pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf("lst is long enough\n")
	}
	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {  // VALID: a *List can be dereferenced for the receiver
		fmt.Printf("plst is long enough\n") // - plst2 is long enough
	}
}
