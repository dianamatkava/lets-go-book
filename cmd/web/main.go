package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func loggingMiddleware(logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("REQ", "Method", r.Method, "URL", r.URL.Path)
		next(w, r) // call actual handler
		logger.Info("RES", "Method", r.Method, "URL", r.URL.Path)
	}
}


func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	address := flag.String("adr", ":4000", "HTTP network address")

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", loggingMiddleware(logger, home))  // Restrict this route to exact matches on / only
	// mux.HandleFunc("/static/", home)  // subtree path pattern, means /static/**, the first matching handler will run
	mux.HandleFunc("GET /snippet", loggingMiddleware(logger, getSnippets))
	mux.HandleFunc("GET /snippet/{id}", loggingMiddleware(logger, getSnippet))
	mux.HandleFunc("POST /snippet", loggingMiddleware(logger, createSnippet))
	mux.HandleFunc("GET /snippet/createForm", loggingMiddleware(logger, getSnippetCreateForm))

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

	logger.Info("starting server on", "address", *address)

	err := http.ListenAndServe(*address, mux)  // how this actually works under the hood?
	logger.Error(err.Error())
	os.Exit(1)
}
