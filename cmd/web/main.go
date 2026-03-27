package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// application-wide dependencies
type application struct {
	logger *slog.Logger
}


func (app *application) loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("REQ", "Method", r.Method, "URL", r.URL.Path)
		next(w, r) // call actual handler
		app.logger.Info("RES", "Method", r.Method, "URL", r.URL.Path)
	}
}


func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := &application{logger: logger}

	address := flag.String("adr", ":4000", "HTTP network address")

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/{$}", app.loggingMiddleware(app.home))  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("GET /snippet", app.loggingMiddleware(app.getSnippets))
	mux.HandleFunc("GET /snippet/{id}", app.loggingMiddleware(app.getSnippet))
	mux.HandleFunc("POST /snippet", app.loggingMiddleware(app.createSnippet))
	mux.HandleFunc("GET /snippet/createForm", app.loggingMiddleware(app.getSnippetCreateForm))
	
	logger.Info("starting server on", "address", *address)

	err := http.ListenAndServe(*address, mux)  // how this actually works under the hood?
	logger.Error(err.Error())
	os.Exit(1)
}
