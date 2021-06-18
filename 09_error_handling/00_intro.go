package main

// For Error Handling Go prefers
// 1- defer-panic-and-recover mechanism
// 2- The Go way to handle errors is for functions and methods to return an error object as their only or last return value—or nil if no error occurred—and for the code calling functions to always check the error they receive.
// 3- Go makes a distinction between critical and non-critical errors: non-critical errors are returned as normal return values, whereas for critical errors, the panic-recover mechanism is used.
// 4- Always assign an error to a variable within a compound if-statement; this makes for clearer code.
// 5- Instead of fmt.Printf, corresponding methods of the log package could be used
// 6- To stop the execution of a program in an error-state, we can use os.Exit(1)
// 7- Whenever you need a new error-type, you can make one with the function errors.New
// 8- fmt.Errorf can be used to make descriptive error on the fly
