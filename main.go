package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := store.db.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()

	server := NewAPIServer(":3000", store)
	server.Run()
}
