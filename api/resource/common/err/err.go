package err

import (
	"fmt"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

type ErrUserID error

func NewErrUserID(e error) ErrUserID {
	return fmt.Errorf("user id error: %w", e)
}

var (
	RespDBDataInsertFailure = []byte(`{"error": "db data insert failure"}`)
	RespDBDataAccessFailure = []byte(`{"error": "db data access failure"}`)
	RespDBDataUpdateFailure = []byte(`{"error": "db data update failure"}`)
	RespDBDataRemoveFailure = []byte(`{"error": "db data remove failure"}`)

	RespJSONEncodeFailure = []byte(`{"error": "json encode failure"}`)
	RespJSONDecodeFailure = []byte(`{"error": "json decode failure"}`)

	RespInvalidURLParamID = []byte(`{"error": "invalid url param-id"}`)
	RespInvalidUserID     = []byte(`{"error": "invalid or missing user id"}`)
	RespResourceNotFound  = []byte(`{"error": "resource not found"}`)
)

func ServerError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}

func BadRequest(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func NotFound(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(resp)
}

func ValidationErrors(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}
