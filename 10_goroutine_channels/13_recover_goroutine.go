package main

import (
	"log"
)

type Work int

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("work failed with %s in %v:", err, work)
			// the failing goroutine is stopped
		}
	}()
	// ...	do(work)
}
