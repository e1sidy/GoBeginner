package main

type Any interface{}
type Element interface{}

// Example 1
// Array of general type can be used to store any type of value
type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, elm Element) {
	p.a[i] = elm
}

// Example 2
// Copying a data-slice in a slice of interface{}

func FuncReturnSlice() []string {
	arr := make([]string, 3)
	return arr
}

// Compilation Error:
// var dataSlice []string = FuncReturnSlice()
// var interfaceSlice []interface{} = dataSlice

// For copying slice to generic slice
// copy by each element manually
func example2() {
	var dataSlice []string = FuncReturnSlice()
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	for ix, d := range dataSlice {
		interfaceSlice[ix] = d
	}
}

// Example 3 Node structures of general or different types
type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}
