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
	address := flag.String("adr", ":4000", "HTTP network address")

	app := &application{logger: logger}
	
	logger.Info("starting server on", "address", *address)

	err := http.ListenAndServe(*address, app.routes())  // how this actually works under the hood?
	logger.Error(err.Error())
	os.Exit(1)
}
