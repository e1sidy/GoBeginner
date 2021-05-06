package main

import (
	"fmt"
	"io"
	"log"
)

/* The defer allows us to guarantee that certain clean-up tasks are performed before we return from a function,
* for example:
* Closing a file stream
* Unlocking a locked resource (a mutex)
* Printing a footer in a report
* Closing a database connection
 */

func main() {
	function1()
	funcLoopDefer()
	traceWithDefer()
	logWithDefer("GO defer")
}

// Differ is used to postpone the execution of a statement or a fucntion
// till the of execution of enclosing(calling) function

func function1() {
	fmt.Println("function1 execution begin")
	defer function2()
	fmt.Println("function1 execution ends")
}

func function2() {
	fmt.Println("function2 is exected after function 1")
}

// After the end of function defer are called in a LIFO manner(stack)
func funcLoopDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Tracing the execution stack using differ
func trace(s string) {
	fmt.Printf("Entering %s\n", s)
}
func untrace(s string) {
	fmt.Printf("Leaving %s\n", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Printf("fucntion a doing something\n")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Printf("function b doing something\n")
	a()
}

func traceWithDefer() {
	b()
}

// Loging with defer
func logWithDefer(s string) (a int, err error) {

	defer func() {
		log.Printf("func1(%s) = %d, %v", s, a, err)
	}()

	return 8, io.EOF

}
