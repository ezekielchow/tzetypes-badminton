// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package userstoregenerated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Player struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Session struct {
	ID                    pgtype.UUID
	UserID                pgtype.UUID
	SessionToken          pgtype.UUID
	RefreshToken          pgtype.UUID
	SessionTokenExpiresAt pgtype.Timestamp
	RefreshTokenExpiresAt pgtype.Timestamp
	CreatedAt             pgtype.Timestamp
	UpdatedAt             pgtype.Timestamp
}

type User struct {
	ID           pgtype.UUID
	Email        string
	PasswordHash *string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}
