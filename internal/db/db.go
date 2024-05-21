package db

import (
	"database/sql"
	"fmt"

	"github.com/agastya909/go-notes/internal/env"
	_ "github.com/go-sql-driver/mysql"
)

func NewSql() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DB_USER, env.DB_PW, env.DB_HOST, env.DB_PORT, env.DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("DB connection established")
	return db, nil
}
