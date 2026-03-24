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
