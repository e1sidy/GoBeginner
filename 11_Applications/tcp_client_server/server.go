package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting the server ...")

	// create listner
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln("Error listening", err.Error())
		os.Exit(1)
	}

	// listen and accept connections from clients:
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("No connection found error:", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		if _, err := conn.Read(buf); err != nil {
			log.Println("Error reading :", err.Error())
			return
		} else {
			fmt.Printf("Received data: %v", string(buf))
		}

	}
}
