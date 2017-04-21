package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"id": 1, "page": "some page"}`))
}
