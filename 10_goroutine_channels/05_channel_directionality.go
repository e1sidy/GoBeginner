package main

/*
- All created channels are bidirectional
- Snipet for unidirectional channels are:-
		var send_only chan<- int 		// data can only be sent (written) to the channel
		var recv_only <-chan int 		// data can only be received (read) from the channel
- Assigning bidirectional channels to directional channel variables snippet
		var c = make(chan int) // bidirectional
		go source(c)
		go sink(c)
		func source(ch chan<- int) {
		for { ch <- 1 } // sending data to ch channel
		}
		func sink(ch <-chan int) {
		for { <-ch } // receiving data from ch channel
		}
*/
