package utils

import (
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
