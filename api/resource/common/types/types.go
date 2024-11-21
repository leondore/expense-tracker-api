package types

import (
	"database/sql"
	"time"

	"golang.org/x/exp/constraints"
)

func NullableStringtoPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func PtrToNullableString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NullableTimeToPtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func PtrToNullableTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}

func NullableInt16ToPtr(i sql.NullInt16) *int16 {
	if i.Valid {
		return &i.Int16
	}
	return nil
}

func PtrToNullableInt16[T constraints.Integer](i *T) sql.NullInt16 {
	if i == nil {
		return sql.NullInt16{}
	}

	return sql.NullInt16{
		Int16: int16(*i),
		Valid: true,
	}
}
