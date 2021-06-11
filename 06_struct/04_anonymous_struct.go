package main

import (
	"fmt"
)

func main() {

	// Two ways of defining an anonymous struct

	var person struct {
		firstName string
		lastName  string
	}

	anotherPerson := struct {
		firstName string
		lastName  string
	}{"Donald", "Trump"}

	person.firstName = "Harry"
	person.lastName = "Putter"

	fmt.Println(person)
	fmt.Println(anotherPerson)
}
