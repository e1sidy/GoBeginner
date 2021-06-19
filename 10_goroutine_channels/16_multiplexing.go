package main

import (
	"fmt"
)

type Request struct {
	a, b   int
	replyc chan int // Reply channel inside Request
}

type binOp func(int, int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit: // clean shutdown of server
			return
		}
	}
}

func startServer(op binOp) (chan *Request, chan bool) {
	service := make(chan *Request)
	quit := make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer(func(i1, i2 int) int { return i1 + i2 })
	const N = 100
	var reqs [N]Request
	for i := 0; i <= N-1; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	// Can be in any order, even random
	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != 2*i+N {
			fmt.Println("Fail at", i)
		} else {
			fmt.Println("Request", i, "is ok!")
		}
	}

	quit <- true
	fmt.Println("done")
}
