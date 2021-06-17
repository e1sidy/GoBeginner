package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(reader *bufio.Reader) {
	for {
		buf, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", buf)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, err := os.Open(flag.Arg(i))

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}

		cat(bufio.NewReader(fd))
	}
}
