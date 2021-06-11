package main

import (
	"fmt"
	"strings"
)

/*
type identifier struct {
  field1 type1
  field2 type2
  ...
}
*/

type struct1 struct {
	num  int
	deci float32
	name string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func example1() {
	obj := new(struct1)

	obj.num = 117
	obj.deci = 22.19
	obj.name = "Seirra"

	fmt.Printf("The fields of object obj are:-\n")
	fmt.Printf("num = %d\n", obj.num)
	fmt.Printf("deci = %f\n", obj.deci)
	fmt.Printf("name = %s\n", obj.name)
	fmt.Println(obj)
}

func example2() {
	// declare struct as value
	var per1 Person
	per1.firstName = "harry"
	per1.lastName = "Osborn"
	upPerson(&per1)
	fmt.Printf("Name = %s %s\n", per1.firstName, per1.lastName)

	// declare struct as pointer
	per2 := new(Person) // Equivalent to per2 := &Person{}
	per2.firstName = "James"
	per2.lastName = "Potter"
	upPerson(per2)
	fmt.Printf("Name = %s %s\n", per2.firstName, per2.lastName)

	// declare struct as a literal
	per3 := &Person{"Chris", "Redfield"}
	upPerson(per3)
	fmt.Printf("Name = %s %s\n", per3.firstName, per3.lastName)

}

func main() {
	example1()
	example2()
}
