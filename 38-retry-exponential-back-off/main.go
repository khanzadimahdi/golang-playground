package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	const timeout = 1 * time.Minute

	fmt.Println(fetch("http://google.com", timeout))
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func fetch(url string, timeout time.Duration) (response *http.Response, err error) {
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		response, err = http.Head(url)

		if err == nil {
			return
		}

		log.Printf("server not responding(%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}

	return nil, fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
