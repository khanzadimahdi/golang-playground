package main

func main() {
	multimirrorQueue()
}

func multimirrorQueue() []byte {
	response := make(chan []byte)

	response <- request("mirror1.com")
	response <- request("mirror2.net")
	response <- request("mirror3.io")

	//----------------------------------------------------------------------------------------------------------
	// the fastest mirror's response would return and the other goroutines would leak.
	// Each remaining request goroutine will block forever when it tries to send a value
	// on response channel, and will never terminate.
	//----------------------------------------------------------------------------------------------------------
	// This situation, a goroutine leak, may cause the whole program to get stuck or to run out of memory.
	//----------------------------------------------------------------------------------------------------------
	// The simplest solution is to use a buffered response channel with sufficient capacity that no request
	// goroutine will block when it sends a message.
	// An alternative solution is to create another goroutine to drain the channel while the main goroutine
	// returns the first error without delay.
	//----------------------------------------------------------------------------------------------------------

	return <-response
}

func request(url string) []byte {
	// send request and return response
	return []byte{
		// some data
	}
}
