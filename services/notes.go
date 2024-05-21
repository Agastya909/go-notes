package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/agastya909/go-notes/internal/types"
	"github.com/agastya909/go-notes/internal/utils"
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
		CreatedOn: time.Now().Local().Add(time.Hour*5 + time.Minute*30),
	}, nil
}

func (s *Store) DeleteById(id string) error {
	row := s.db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return errors.New(utils.MESSAGES["NOT_ID_INVALID"])
	}
	_, err = s.db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return errors.New(utils.MESSAGES["COULD_NOT_DELETE"])
	}
	return nil
}

func (s *Store) UpdateById(id string) error {
	return nil
}

func (s *Store) Save(NewNote types.NewNote) error {
	id, createdOn := utils.GetUUID(), time.Now().Local().Add(time.Hour*5+time.Minute*30)
	_, err := s.db.Exec("INSERT INTO notes (id, title, body, created_on) VALUES(?,?,?,?)", id, NewNote.Title, NewNote.Body, createdOn)
	if err != nil {
		return errors.New(utils.MESSAGES["COULD_NOT_SAVE"])
	}
	return nil
}
