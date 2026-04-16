package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/dianamatkava/snippetbox/cmd/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// application-wide dependencies
type application struct {
	logger 		*slog.Logger
	snippets 	*models.SnippetModel
	templates 	map[string]*template.Template
}


func openDB(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}


func main() {
	address := flag.String("adr", ":4000", "HTTP network address")
	
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	
	db, dbErr := openDB("pgx", "postgres://postgres:postgres@localhost/postgres")
	if dbErr != nil {
		logger.Error(dbErr.Error())
		os.Exit(1)
	}
	defer db.Close()
	

	templateCache, err := cacheTemplates()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: 	logger,
		snippets: 	&models.SnippetModel{DB: db},
		templates: 	templateCache,
	}
	
	logger.Info("starting server on", "address", *address)

	err = http.ListenAndServe(*address, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
