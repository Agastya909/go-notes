package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewSql() (*sql.DB, error) {
	var (
		DB_USER = os.Getenv("DB_USER")
		DB_PW   = os.Getenv("DB_PW")
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", DB_USER, DB_PW, DB_HOST, DB_PORT)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = InitDB(db)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("DB connection established")
	return db, nil
}

func InitDB(db *sql.DB) error {
	// create db
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS notes")
	if err != nil {
		return fmt.Errorf("could not create db: %v", err.Error())
	}
	// use db
	_, err = db.Exec("USE notes")
	if err != nil {
		return fmt.Errorf("could not select notes as db: %v", err.Error())
	}
	// create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS notes (
  		id VARCHAR(255) PRIMARY KEY,
    	title VARCHAR(255),
    	body TEXT,
    	created_on DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_on DATETIME DEFAULT NULL
		)`)
	if err != nil {
		return fmt.Errorf("could not create table: %v", err.Error())
	}
	return nil
}
