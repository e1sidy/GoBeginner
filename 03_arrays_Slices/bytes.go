package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buff bytes.Buffer

	// Writing to a buffer using bytes.Buffer.WriteString(str)
	buff.WriteString("ABC")
	buff.WriteString("\teLSEIDy\t")

	// Writing to buffer using fmt.Fprintf()
	fmt.Fprintf(&buff, " Number of days to incident %d", 5)

	// Printing the buffer
	fmt.Println(buff.String())
}
