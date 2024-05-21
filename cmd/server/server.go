package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/agastya909/go-notes/cmd/routes"
	"github.com/agastya909/go-notes/services"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
}

func New(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	subRouter := r.NewRoute().PathPrefix("/api/v1/").Subrouter()

	InitNotesRoute(subRouter, s)

	log.Println("server listening on port", s.addr)
	return http.ListenAndServe(s.addr, r)
}

func InitNotesRoute(r *mux.Router, s *Server) {
	r = r.NewRoute().PathPrefix("/notes").Subrouter()
	noteStore := services.NewStore(s.db)
	noteService := routes.NoteHandler(noteStore)
	noteService.NoteRoutes(r)
}
