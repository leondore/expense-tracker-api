package err

import (
	"fmt"
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
)
