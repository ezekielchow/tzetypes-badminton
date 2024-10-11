// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package clubstoregenerated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createClub = `-- name: CreateClub :one
INSERT INTO clubs (
  owner_id, name
) VALUES (
  $1,$2
) RETURNING id, owner_id, name, created_at, updated_at
`

type CreateClubParams struct {
	OwnerID pgtype.UUID
	Name    string
}

func (q *Queries) CreateClub(ctx context.Context, arg CreateClubParams) (Club, error) {
	row := q.db.QueryRow(ctx, createClub, arg.OwnerID, arg.Name)
	var i Club
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
