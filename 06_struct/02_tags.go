package main

// Tags are strings attached to a field in a struct
// It can be used for documentation or some imp label

import (
	"fmt"
	"reflect"
)

type T struct {
	a int "This is a tag"
	b int `A raw string tag`
	c int `key1:"value1" key2:"value2"`
}

func refTag(obj T, idx int) {
	structType := reflect.TypeOf(obj)
	ixField := structType.Field(idx) // getting field at a position ix
	fmt.Printf("%v\n", ixField.Tag)  // printing tags
}

func example1() {
	obj := T{1, 4, 5}
	for i := 0; i < 3; i++ {
		refTag(obj, i)
	}
}

// Different ways of refrencing a tag
// 1- 	reflect.TypeOf(obj).Field(0).Tag
// 2- 	field, ok := reflect.TypeOf(obj).FieldByName("b")
//		if ok {field.Tag}
// 3- 	field.Tag.Get("key1")

func example2() {
	obj := T{}
	fmt.Println(reflect.TypeOf(obj).Field(0).Tag)
	if field, ok := reflect.TypeOf(obj).FieldByName("b"); ok {
		fmt.Println(field.Tag)
	}
	if field, ok := reflect.TypeOf(obj).FieldByName("c"); ok {
		fmt.Println(field.Tag)
		fmt.Println(field.Tag.Get("key1"))
	}
	if field, ok := reflect.TypeOf(obj).FieldByName("d"); ok {
		fmt.Println(field.Tag)
	} else {
		fmt.Println("Field not found")
	}
}

func main() {
	example1()
	fmt.Println()
	example2()
}
