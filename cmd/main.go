package main

import (
	"log"

	"github.com/agastya909/go-notes/cmd/server"
	"github.com/agastya909/go-notes/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load dotenv file")
	}
	// start db
	db, err := db.NewSql()
	if err != nil {
		log.Fatal("could not connect to db: ", err.Error())
	}
	// init server
	server := server.New(":8080", db)
	if err := server.Start(); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
