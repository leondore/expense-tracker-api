package types

import (
	"database/sql"
	"time"
)

func NullableStringtoPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func NullableTimeToPtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
