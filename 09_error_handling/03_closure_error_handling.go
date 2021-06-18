package main

/*
Use of closure for error handling
 - Whenever a function returns, we should test whether it results in an error: this can lead to repetitive code.
 - Combining the defer/panic/recover mechanisms with closures can result in a far more elegant
 - However, it is only applicable when all functions have the same signature, which is rather restrictive.
 - A good example of its use is in web applications, where all handler functions are of the following type:
	func handler1(w http.ResponseWriter, r *http.Request) { ... }
*/

// Example
// Suppose all functions have the signature: func f(a type1, b type2)
// The number of parameters and their types is irrelevant

import (
	"log"
)

type type1 int32
type type2 string
type fType1 func(a type1, b type2)

// check: a function which tests whether an error occurred, and panics if so:
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// errorhandler: this is a wrapper function. It takes a function fn of our type fType1 as parameter, and returns such a function by calling fn. However, it contains the defer/recover mechanism:
func errorHandler(fn fType1) fType1 {
	return func(a type1, b type2) {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		fn(a, b)
	}
}

// The check() function is used in every called function, like this:
/*
func f1(a type1, b type2) {
	...
	f, _, err := // call function/method
	check(err)
	t, err = // call function/method
	check(err)
	_, err = // call function/method
	check(err)
	...
}
*/

//The main() or other caller-function should then call the necessary functions wrapped in errorHandler, like this:
/*
func main() {
  errorHandler(f1)
  errorHandler(f2)
  ...
}
*/
