package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func example1() {
	inputFile, inputError := os.Open("./input_files/input.dat")
	if inputError != nil {
		fmt.Println("Error occured in opening the file")
		return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input is: %s", inputString)
	}
}

// Reading entire content of a file in a string
func example2() {

	// buf is of type []byte
	// we might need to convert in to string by string(buf)

	//Don’t use ReadFile for big files because one large string uses a lot of memory!

	buf, err := ioutil.ReadFile("./input_files/product.txt")
	if err != nil {
		fmt.Println("Error in opening file")
		return
	}

	fmt.Println(string(buf))
}

// Reading columns of data from a file
func example3() {
	file, err := os.Open("./input_files/products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func example4() {
	fName := "./input_files/Example.json.gz"
	var reader *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)

	if err != nil {
		reader = bufio.NewReader(fi)
	} else {
		reader = bufio.NewReader(fz)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		} else {
			fmt.Println(line)
		}
	}
}

func main() {
	example1()
	example2()
	example3()
	example4()
}
