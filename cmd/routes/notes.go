package routes

import (
	"fmt"
	"net/http"

	"github.com/agastya909/go-notes/internal/types"
	"github.com/agastya909/go-notes/internal/utils"
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
	r.HandleFunc("/", h.Save).Methods("POST")
	r.HandleFunc("/id", h.GetById).Methods("GET")
	r.HandleFunc("/{id}", h.DeleteById).Methods("DELETE")
	r.HandleFunc("/{id}", h.UpdateById).Methods("PATCH")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := h.note.GetAll()
	if err != nil {
		utils.WriteHttpError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteHttpSuccess(w, http.StatusOK, utils.MESSAGES["NOTE_FOUND"], notes)
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
	err := h.note.DeleteById(id)
	if err != nil {
		utils.WriteHttpError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteHttpSuccess(w, http.StatusOK, utils.MESSAGES["NOTE_DELETED"], nil)
}

func (h *handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	h.note.UpdateById(id)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Save(w http.ResponseWriter, r *http.Request) {
	var payload types.NewNote
	err := utils.ParseJsonRequest(r, &payload)
	if err != nil {
		utils.WriteHttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.note.Save(payload)
	if err != nil {
		utils.WriteHttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteHttpSuccess(w, http.StatusOK, utils.MESSAGES["NOTE_SAVED"], nil)
}
