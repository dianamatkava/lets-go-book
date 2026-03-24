package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}
	file, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = file.ExecuteTemplate(w, "base", nil)  // write headers and body
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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
