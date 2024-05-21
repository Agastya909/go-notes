package types

import (
	"time"
)

type Note struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

type NewNote struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteRepository interface {
	GetAll() ([]Note, error)
	GetById(id string) (Note, error)
	DeleteById(id string) error
	UpdateById(id string) error
	Save(NewNote) error
}
