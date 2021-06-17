package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func example1() {
	// Converting a JSON format byte into a struct
	b := []byte(`{"Value": "Father", "Left": {"Value": "Left child"}, "Right": {"Value": "Right
	child"}}`)
	var father Node
	json.Unmarshal(b, &father)
	fmt.Println(father, father.Left, father.Right)
}

// json and struct can have different field name
// struct field tags must be used in that case
type Person struct {
	Name string `json:"personName"`
	Age  int    `json:"personAge"`
}

func example2() {
	b := []byte(`{"personName": "Obama", "personAge": 57}`)

	var person Person
	json.Unmarshal(b, &person)
	fmt.Println(person)

	js, _ := json.Marshal(person)
	fmt.Println(string(js))
}

func main() {
	example1()
	example2()
}
