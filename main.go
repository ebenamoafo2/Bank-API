package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	store, err := NewPostgresStore()
	if err != nil {
		slog.Error("failed to create store", "error", err)
		os.Exit(1)
	}

	if err := store.Init(); err != nil {
		slog.Error("failed to initialize store", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := store.db.Close(); err != nil {
			slog.Error("error closing database", "error", err)
		}
	}()

	server := NewAPIServer(":3000", store)
	server.Run()
}
