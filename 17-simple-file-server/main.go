package main

import (
	"fmt"
	"net/http"
)

const addr string = "127.0.0.1:8080"

func main() {
	fmt.Println(`static file server is running on: http://` + addr)

	if err := http.ListenAndServe(addr, http.FileServer(http.Dir("./static"))); err != nil {
		fmt.Println(err)
	}
}
