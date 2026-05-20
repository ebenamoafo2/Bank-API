package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func seedAccount(store storage, firstName, lastName, password string) *Account {
	account, err := NewAccount(firstName, lastName, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(account); err != nil {
		log.Fatal("Seeding failed:")
	}

	return account
}

func seedAccounts(store storage) {
	seedAccount(store, "John", "Doe", "password")
	seedAccount(store, "Jane", "Doe", "password")
}

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, relying on system environment variables")
	}

	seed := flag.Bool("seed", false, "seed the database with some accounts")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// load once at startup
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		slog.Error("JWT_SECRET environment variable is not set")
		os.Exit(1)
	}

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

	if *seed {
		fmt.Println("Seeding database...")
		seedAccounts(store)
		os.Exit(0)
	}

	server := NewAPIServer(":3000", store, jwtSecret) // pass it in
	server.Run()
}
