package utils

import (
	"encoding/json"
	"fmt"
	"io"
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
	"NOTE_NOT_FOUND":   "note not found",
	"NOTE_FOUND":       "note found",
	"NOTE_UPDATED":     "note updated",
	"COULD_NOT_UPDATE": "could not update the note",
}

func ParseJsonRequest(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	if len(body) == 0 {
		return fmt.Errorf("missing request body")
	}

	err = json.Unmarshal(body, payload)
	if err != nil {
		return fmt.Errorf("error decoding json: %v", err)
	}

	return nil
}

func writeJsonResponse(w http.ResponseWriter, statusCode int, res any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(res)
}

func WriteHttpError(w http.ResponseWriter, statusCode int, message string) {
	writeJsonResponse(w, statusCode, map[string]string{"message": message})
}

func WriteHttpSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	writeJsonResponse(w, statusCode, map[string]any{"message": message, "data": data})
}

func GetUUID() string {
	return uuid.New().String()
}
