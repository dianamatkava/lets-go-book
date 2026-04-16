package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/{$}", app.home)  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("GET /snippet", app.getSnippets)
	mux.HandleFunc("GET /snippet/{id}", app.getSnippet)
	mux.HandleFunc("POST /snippet", app.createSnippet)
	mux.HandleFunc("GET /snippet/createForm", app.getCreateFormSnippet)

	middlewares := alice.New(app.panicRecoverMiddleware, app.loggingMiddleware, SecureHeadersMiddleware)
	return middlewares.Then(mux)
}