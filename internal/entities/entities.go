package entities

import (
	"database/sql"
)

func NewNullString(s string) *sql.NullString {
	if len(s) == 0 {
		return &sql.NullString{}
	} else {
		return &sql.NullString{
			String: s,
			Valid:  true,
		}
	}
}

type Record struct {
	TickerID  int64
	Timestamp int64
	Price     string
}
