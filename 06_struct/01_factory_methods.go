package main

import (
	"fmt"
)

type File struct {
	fd   int    // file descriptor number
	name string // filename
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	file := NewFile(1, "abc.xyz")
	fmt.Println(file)
}
