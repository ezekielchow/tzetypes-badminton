// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  firebase_uid,email,account_tier
) VALUES (
  $1::text,
  $2::text,
  $3::text
) RETURNING id, firebase_uid, email, account_tier, created_at, updated_at
`

type CreateUserParams struct {
	FirebaseUid string
	Email       string
	AccountTier string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.FirebaseUid, arg.Email, arg.AccountTier)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirebaseUid,
		&i.Email,
		&i.AccountTier,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserWithEmail = `-- name: FindUserWithEmail :one
SELECT id, firebase_uid, email, account_tier, created_at, updated_at FROM users WHERE email=$1
`

func (q *Queries) FindUserWithEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, findUserWithEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirebaseUid,
		&i.Email,
		&i.AccountTier,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserWithID = `-- name: FindUserWithID :one
SELECT id, firebase_uid, email, account_tier, created_at, updated_at FROM users WHERE id=$1
`

func (q *Queries) FindUserWithID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUserWithID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirebaseUid,
		&i.Email,
		&i.AccountTier,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
