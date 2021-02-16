package main

import (
	"fmt"
	"log"
	"net/http"
)

type Database map[string]map[string]interface{}

func main() {
	db := Database{
		"mahdi": {
			"age":     27,
			"country": "Germany",
		},
		"asghar": {
			"age":     17,
			"country": "USA",
		},
	}

	mux := http.NewServeMux()

	// handle functions can be added using one of the below syntaxes
	mux.HandleFunc("/", db.Show)
	mux.Handle("/names", http.HandlerFunc(db.ShowNames))

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (d Database) Show(w http.ResponseWriter, rq *http.Request) {
	for name, item := range d {
		fmt.Fprintf(w, "%s \n", name)
		for key, value := range item {
			fmt.Fprintf(w, "\t%s = %v\n", key, value)
		}
	}
}

func (d Database) ShowNames(w http.ResponseWriter, rq *http.Request) {
	for name, _ := range d {
		fmt.Fprintf(w, "%s\n", name)
	}
}
