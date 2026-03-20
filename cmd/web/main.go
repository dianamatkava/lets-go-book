package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQ %s %s", r.Method, r.URL.Path)
		next(w, r) // call actual handler
		log.Printf("RES %s %s", r.Method, r.URL.Path)
	}
}


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getSnippets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display all snippets")
}

func getSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Get specific snippet by ID: %d", id)
}

func getSnippetCreateForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create form")
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.WriteHeader(http.StatusCreated)  // why I dont specify that it is a status code and not something else?
	fmt.Fprintf(w, "Create snippet %d", 12)
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
