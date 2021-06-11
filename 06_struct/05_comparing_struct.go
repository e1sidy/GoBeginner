package main

// Structs can only be compared if they are of same type
// Structs can be converted from one type to another if they have same fields(type)

import (
	"fmt"
)

type T1 struct {
	a int
}

type T2 struct {
	b int
}

type T3 struct {
	a int
}

func main() {
	t1 := T1{10}
	t2 := T2{10}
	t3 := T3{10}
	t4 := T3{20}
	fmt.Println(t1, t2, t3, t4)
	// fmt.Println("t1 == t2?", t1==t2) // <-- invalid operation: t1 == t2 (mismatched types T1 and T2)
	// fmt.Println("t1 == t3?", t1==t3) // <-- invalid operation: t1 == t3 (mismatched types T1 and T3)
	fmt.Println("t1 == t3?", t1 == T1(t3)) // true
	fmt.Println("t3 == t4?", t3 == t4)     // false
}
