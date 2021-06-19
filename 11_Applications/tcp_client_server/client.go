package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Open connection
	conn, err := net.Dial("tcp", ":3001")
	if err != nil {
		// No connection could be made because the target machine actively refused it.
		log.Println("Error dialing", err.Error())
		os.Exit(1) // terminate program
	}

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your name?")
	clientName, _ := inputReader.ReadString('\n')

	trimmedClient := strings.Trim(clientName, "\n")

	// send info to server until Quit:
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		if _, err := conn.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n")); err != nil {
			log.Println("Error in sending data to server Error:", err.Error())
		}
	}
}
