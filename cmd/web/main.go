package main

import (
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


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", loggingMiddleware(home))  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("GET /snippet", loggingMiddleware(getSnippets))
	mux.HandleFunc("GET /snippet/{id}", loggingMiddleware(getSnippet))
	mux.HandleFunc("POST /snippet", loggingMiddleware(createSnippet))
	mux.HandleFunc("GET /snippet/createForm", loggingMiddleware(getSnippetCreateForm))

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)  // how this actually works under the hood?
	log.Fatal(err) // how to do other log levels? is it used in prod apps?
}
