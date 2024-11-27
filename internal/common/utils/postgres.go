package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringToPgId(id string) (pgtype.UUID, error) {
	pgID := pgtype.UUID{}
	err := pgID.Scan(id)
	if err != nil {
		return pgID, err
	}

	return pgID, err
}

func TimeToPgTimestamp(time time.Time) (pgtype.Timestamp, error) {
	pgTimestamp := pgtype.Timestamp{}
	err := pgTimestamp.Scan(time)
	if err != nil {
		return pgTimestamp, err
	}

	return pgTimestamp, nil
}
