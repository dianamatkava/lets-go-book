package main

import (
	"log/slog"
	"os"
	"log"
)


func main() {
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(loggerHandler)

	log.Fatal("Error")

	logger.Info("Hello World")
}