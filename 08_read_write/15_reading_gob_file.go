package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {

	var vc VCard
	file, err := os.Open("output/vcard.gob")

	if err != nil {
		log.Fatal("File not found")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	dec := gob.NewDecoder(reader)
	err = dec.Decode(&vc)
	if err != nil {
		log.Fatal("Error in decoding")
		os.Exit(1)
	}
	fmt.Println("Received Data : ", vc)

}
