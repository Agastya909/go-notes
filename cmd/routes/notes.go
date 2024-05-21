package routes

import (
	"fmt"
	"net/http"

	"github.com/agastya909/go-notes/internal/types"
	"github.com/gorilla/mux"
)

type handler struct {
	note types.NoteRepository
}

func NoteHandler(note types.NoteRepository) *handler {
	return &handler{note: note}
}

func (h *handler) NoteRoutes(r *mux.Router) {
	r.HandleFunc("/", h.GetAll).Methods("GET")
	r.HandleFunc("/id", h.GetById).Methods("GET")
	r.HandleFunc("/{id}", h.DeleteById).Methods("DELETE")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	h.note.GetAll()
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	res, err := h.note.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(res)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	h.note.DeleteById(id)
	w.WriteHeader(http.StatusOK)
}
