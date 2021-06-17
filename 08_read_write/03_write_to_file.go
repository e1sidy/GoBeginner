package main

import (
	"bufio"
	"fmt"
	"os"
)

func example1() {
	outputFile, err := os.OpenFile("output/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	outputString := "Hello World from code to file!\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(outputString)
	}
	writer.Flush()
}

func example2() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("output/test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()

	f.WriteString("hello, world in a file\n")
}

func main() {
	example1()
	example2()
}
