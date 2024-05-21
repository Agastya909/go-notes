package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var MESSAGES = map[string]string{
	"NO_NOTE_BODY":     "empty or incomplete note, add both title and body",
	"COULD_NOT_SAVE":   "could not save the note",
	"NOTE_SAVED":       "Note saved succesfully",
	"NOTE_ID_REQUIRED": "note id is required",
	"COULD_NOT_DELETE": "could not delete the note",
	"NOTE_DELETED":     "Note deleted succesfully",
	"NOT_ID_INVALID":   "note with this id not found",
}

func ParseJsonRequest(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, res any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(res)
}

func WriteHttpError(w http.ResponseWriter, statusCode int, message string) {
	WriteJsonResponse(w, statusCode, map[string]string{"message": message})
}

func WriteHttpSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	WriteJsonResponse(w, statusCode, map[string]any{"message": message, "data": data})
}

func GetUUID() string {
	return uuid.New().String()
}
