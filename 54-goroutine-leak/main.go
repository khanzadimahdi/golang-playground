package main

func main() {
	multimirrorQueue()
}

func multimirrorQueue() []byte {
	response := make(chan []byte)

	response <- request("mirror1.com")
	response <- request("mirror2.net")
	response <- request("mirror3.io")

	// the fastest mirror's response would return and the other goroutines would leak.
	return <-response
}

func request(url string) []byte {
	// send request and return response
	return []byte{
		// some data
	}
}
