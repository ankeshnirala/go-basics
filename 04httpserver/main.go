package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string
	Author string
	Pages  int16
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 color=red>Hello</h1>"))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := Book{Title: "Atomic Habbit", Author: "Ankesh Kumar", Pages: 405}

	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/book", bookHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
