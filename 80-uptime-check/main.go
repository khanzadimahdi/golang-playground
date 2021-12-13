package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s is down: %s\n , status: %d\n", url, err, resp.StatusCode)
	} else {
		defer resp.Body.Close()

		fmt.Printf("%s is up\n", url)
	}

	c <- url
}

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"http://ditty.ir",
		"https://roocket.ir",
		"http://tarhche.ir",
	}

	c := make(chan string)

	fmt.Println("starting to check uptime:")

	fmt.Println(strings.Repeat("-", 40))

	for _, url := range urls {
		go func(url string) {
			c <- url
		}(url)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(time.Second * 2)

			checkURL(url, c)
		}(url)
	}
}
