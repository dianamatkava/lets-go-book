package main

import (
	"fmt"
	"log"
	"net/http"
)


func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQ %s %s", r.Method, r.URL.Path)
		next(w, r) // call actual handler
		log.Printf("RES %s %s", r.Method, r.URL.Path)
	}
}


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", loggingMiddleware(home))  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("/snippet/view", loggingMiddleware(snippetView))
	mux.HandleFunc("/snippet/create", loggingMiddleware(snippetCreate))

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)  // how this actually works under the hood?
	log.Fatal(err) // how to do other log levels? is it used in prod apps?
}
