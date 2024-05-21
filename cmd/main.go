package main

import (
	"log"

	"github.com/agastya909/go-notes/cmd/server"
	"github.com/agastya909/go-notes/internal/db"
)

func main() {
	// init db
	db, err := db.NewSql()
	if err != nil {
		log.Fatal("could not connect to db")
	}
	// init server
	server := server.New(":8080", db)
	if err := server.Start(); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
