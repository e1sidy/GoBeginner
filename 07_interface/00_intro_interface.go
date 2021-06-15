/*
Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.

The name of an interface is formed by the method name plus the [er] suffix. If [er] is not good suffix then use [able]
:-
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}


Unlike OO Programming Language interfaces in GO can have variables of type interface
:-
var ai Namer

ai is a multiword data structure with an uninitialized value of nil. Although not entirely the same thing, ai is, in essence, a pointer. So, pointers to interface values are illegal; they would be wholly useless and give rise to errors in code.


Note the following important points:
1- A type doesn’t have to state explicitly that it implements an interface; interfaces are satisfied implicitly.
2- Multiple types can implement the same interface.
3- A type that implements an interface can also have other functions.
4- A type can implement many interfaces.
5- An interface type can contain a reference to an instance of any of the types that implement the interface.
*/

package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectange struct {
	length, width float32
}

func (this *Square) Area() float32 {
	return this.side * this.side
}

func (this Rectange) Area() float32 {
	return this.length * this.width
}

func example1() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	// Here Square does not explicitly state it implements Shaper
	// yet variable of type Shaper can refer to variable of type Square
	areaIntf = sq1
	// shorter, without separate declaration:
	// areaIntf := Shaper(sq1)

	fmt.Printf("The square has an area of %.2f\n", areaIntf.Area())
}

func example2() {
	sq1 := &Square{5}
	rec1 := Rectange{8, 9}

	// If a struct method (which is in interface method) is defined over value
	// One cannot use interface to refer to  struct varible initalized as a pointer
	// Eg Rectange can not be used as pointer

	shapes := []Shaper{sq1, rec1}
	fmt.Println("Looping throught all the shapes")
	for idx, val := range shapes {
		fmt.Printf("The Area of Shape at index %d is %f\n", idx, val.Area())
	}

}

func main() {
	fmt.Println("Example 1")
	example1()
	fmt.Println("Example 2")
	example2()
}
