package main

import (
	"fmt"
	"log"
	"net/http"
)

type Database map[string]string

func main() {
	db := Database{
		"name":   "mahdi",
		"family": "khanzadi",
	}

	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func (d Database) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	for index, item := range d {
		fmt.Fprintf(w, "%s = %s \n", index, item)
	}
}
