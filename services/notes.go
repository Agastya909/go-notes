package services

import (
	"database/sql"
	"errors"
	"fmt"
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
	rows, err := s.db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, errors.New(utils.MESSAGES["NOTE_NOT_FOUND"])
	}
	defer rows.Close()
	var notes []types.Note
	for rows.Next() {
		var note types.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.CreatedOn)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (s *Store) GetById(id string) (types.Note, error) {
	row := s.db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
	var note types.Note
	err := row.Scan(&note.Id, &note.Title, &note.Body, &note.CreatedOn)
	if err != nil {
		return types.Note{}, fmt.Errorf(utils.MESSAGES["NOT_ID_INVALID"])
	}
	return note, nil
}

func (s *Store) DeleteById(id string) error {
	row := s.db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return fmt.Errorf(utils.MESSAGES["NOT_ID_INVALID"])
	}

	_, err = s.db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf(utils.MESSAGES["COULD_NOT_DELETE"])
	}
	return nil
}

func (s *Store) UpdateById(id string, payload types.NewNote) error {
	if len(payload.Title) == 0 || len(payload.Body) == 0 {
		return fmt.Errorf(utils.MESSAGES["NO_NOTE_BODY"])
	}
	currentTime := time.Now().Local().Add(time.Hour*5 + time.Minute*30)

	_, err := s.db.Exec("UPDATE notes SET title = ?, body = ?, updated_on = ? WHERE id = ? ", payload.Title, payload.Body, currentTime, id)
	if err == sql.ErrNoRows {
		return fmt.Errorf(utils.MESSAGES["NOT_ID_INVALID"])
	}
	return nil
}

func (s *Store) Save(NewNote types.NewNote) error {
	id, createdOn := utils.GetUUID(), time.Now().Local().Add(time.Hour*5+time.Minute*30)
	if len(NewNote.Body) == 0 || len(NewNote.Title) == 0 {
		return fmt.Errorf(utils.MESSAGES["NO_NOTE_BODY"])
	}

	_, err := s.db.Exec("INSERT INTO notes (id, title, body, created_on) VALUES(?,?,?,?)", id, NewNote.Title, NewNote.Body, createdOn)
	if err != nil {
		return fmt.Errorf(utils.MESSAGES["COULD_NOT_SAVE"])
	}
	return nil
}
