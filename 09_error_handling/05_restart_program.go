package main

// This program is just a template
// The following code snippet recovers from the panic and restarts the program

import (
	"log"
	"os"
	"os/exec"
)

func ultraCrazyFunction() {
	prog := os.Args[0]
	if e := recover(); e != nil {
		cmd := exec.Command(prog)
		err := cmd.Run()

		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func main() {
	defer ultraCrazyFunction()
}
