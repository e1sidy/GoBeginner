package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	books := make([]Book, 1)
	fd, err := os.Open("./input_files/products.txt")

	if err != nil {
		log.Fatalf("Error %s opening file products.txt", err)
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error", err)
		}

		line = string(line[:len(line)-1])

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]

		if book.price, err = strconv.ParseFloat(strSl[1], 32); err != nil {
			fmt.Println("Error in file data", err)
		}

		if book.quantity, err = strconv.Atoi(strSl[2]); err != nil {
			fmt.Println("Error in file data", err)
		}

		if books[0].title == "" {
			books[0] = *book
		} else {
			books = append(books, *book)
		}
	}
	fmt.Println("Books Data:")
	for _, book := range books {
		fmt.Println(book)
	}
}
