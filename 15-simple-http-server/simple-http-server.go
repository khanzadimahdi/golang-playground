package main

import (
	"fmt"
	"net/http"
)

func httpHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var response string = "<h1>hello world</h1>"

	fmt.Fprint(responseWriter, response)
}

func main() {
	http.HandleFunc("/", httpHandler)

	http.ListenAndServe(":8000", nil)
}