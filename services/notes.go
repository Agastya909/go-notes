package services

import (
	"database/sql"
	"time"

	"github.com/agastya909/go-notes/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAll() ([]types.Note, error) {
	return nil, nil
}

func (s *Store) GetById(id string) (types.Note, error) {
	return types.Note{
		Id:        "j",
		Title:     "t",
		Body:      "b",
		CreatedOn: time.Now(),
	}, nil
}

func (s *Store) DeleteById(id string) error {
	return nil
}
