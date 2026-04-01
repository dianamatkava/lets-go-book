package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"net"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// application-wide dependencies
type application struct {
	logger *slog.Logger
	db *sql.DB
}


func (app *application) loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("REQ", "Method", r.Method, "URL", r.URL.Path)
		next(w, r) // call actual handler
		app.logger.Info("RES", "Method", r.Method, "URL", r.URL.Path)
	}
}

func openDB(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if errPing != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}


func main() {
	net.Listen()
	address := flag.String("adr", ":4000", "HTTP network address")
	
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	
	db, dbErr := openDB("pgx", "postgres://postgres:postgres@localhost/postgres?parseTime=true")
	if dbErr != nil {
		logger.Error(dbErr.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
		db: db,
	}
	
	logger.Info("starting server on", "address", *address)

	err := http.ListenAndServe(*address, app.routes())  // how this actually works under the hood?
	logger.Error(err.Error())
	os.Exit(1)
}
