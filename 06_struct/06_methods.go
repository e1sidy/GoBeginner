package main

import (
	"fmt"
)

// General Format for a method is
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

// Generally we should pass the receiver as a pointer for performance reasons
// Go automaticaly takes care whether to convert the variable to pointer or value for applying a method on it

type TwoInts struct {
	a, b int
}

func (obj TwoInts) addBoth() int {
	return obj.a + obj.b
}

func (obj TwoInts) addParams(param int) {
	obj.a += param
	obj.b += param
}

// Passing the value by reference to make changes to it
func (obj *TwoInts) addParamsRef(param int) {
	obj.a += param
	obj.b += param
}

func main() {

	obj := new(TwoInts)
	obj.a = 4
	obj.b = 6
	fmt.Println(obj.addBoth())
	obj.addParams(10)
	fmt.Println(obj.addBoth())
	// When passing pointer to struct the fields of struct can be changed
	obj.addParamsRef(10)
	fmt.Println(obj.addBoth())
}
