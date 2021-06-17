package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {

	hasher := sha1.New()

	// Ways of hashing
	// 1 - Using io.WriteString
	io.WriteString(hasher, "test")
	// Printing the checksum
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")

	// 2 - Using Hash.Write()
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
