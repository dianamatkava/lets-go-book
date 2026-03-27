package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/{$}", app.loggingMiddleware(app.home))  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("GET /snippet", app.loggingMiddleware(app.getSnippets))
	mux.HandleFunc("GET /snippet/{id}", app.loggingMiddleware(app.getSnippet))
	mux.HandleFunc("POST /snippet", app.loggingMiddleware(app.createSnippet))
	mux.HandleFunc("GET /snippet/createForm", app.loggingMiddleware(app.getSnippetCreateForm))

	return mux
}