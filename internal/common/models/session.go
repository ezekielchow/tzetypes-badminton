package models

import (
	"os"
	"strconv"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Session struct {
	ID                    string
	UserID                string
	SessionToken          string
	RefreshToken          string
	SessionTokenExpiresAt time.Time
	RefreshTokenExpiresAt time.Time
	CreatedAt             time.Time
	UpdatedAt             *time.Time
}

func (s *Session) PostgresToModel(fromDb database.Session) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	userID, err := uuid.FromBytes(fromDb.UserID.Bytes[:])
	if err != nil {
		return err
	}

	sessionToken, err := uuid.FromBytes(fromDb.SessionToken.Bytes[:])
	if err != nil {
		return err
	}

	refreshToken, err := uuid.FromBytes(fromDb.RefreshToken.Bytes[:])
	if err != nil {
		return err
	}

	s.ID = id.String()
	s.UserID = userID.String()
	s.SessionToken = sessionToken.String()
	s.SessionTokenExpiresAt = fromDb.SessionTokenExpiresAt.Time
	s.RefreshToken = refreshToken.String()
	s.RefreshTokenExpiresAt = fromDb.RefreshTokenExpiresAt.Time
	s.CreatedAt = fromDb.CreatedAt.Time
	s.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}

func (session *Session) ModelToPostgres(model Session) (database.Session, error) {

	s := database.Session{}

	id := pgtype.UUID{}
	err := id.Scan(model.ID)
	if err != nil {
		return s, err
	}

	userID := pgtype.UUID{}
	err = userID.Scan(model.UserID)
	if err != nil {
		return s, err
	}

	sessionToken := pgtype.UUID{}
	err = sessionToken.Scan(model.SessionToken)
	if err != nil {
		return s, err
	}

	sessionTokenExpiry := pgtype.Timestamp{}
	err = sessionTokenExpiry.Scan(model.SessionTokenExpiresAt)
	if err != nil {
		return s, err
	}

	refreshToken := pgtype.UUID{}
	err = refreshToken.Scan(model.RefreshToken)
	if err != nil {
		return s, err
	}

	refreshTokenExpiry := pgtype.Timestamp{}
	err = refreshTokenExpiry.Scan(model.RefreshTokenExpiresAt)
	if err != nil {
		return s, err
	}

	createdAt := pgtype.Timestamp{}
	err = createdAt.Scan(model.CreatedAt)
	if err != nil {
		return s, err
	}

	updatedAt := pgtype.Timestamp{}
	err = updatedAt.Scan(model.UpdatedAt)
	if err != nil {
		return s, err
	}

	s.ID = id
	s.UserID = userID
	s.SessionToken = sessionToken
	s.SessionTokenExpiresAt = sessionTokenExpiry
	s.RefreshToken = refreshToken
	s.RefreshTokenExpiresAt = refreshTokenExpiry
	s.CreatedAt = createdAt
	s.UpdatedAt = updatedAt

	return s, nil
}

func (s Session) NewSessionTokenExpiry() (*time.Time, error) {
	sessionLifespan, err := strconv.Atoi(os.Getenv("SESSION_LIFESPAN_MINUTES"))
	if err != nil {
		return nil, err
	}

	sessionExpiresAt := time.Now().Add(time.Minute * time.Duration(sessionLifespan))
	return &sessionExpiresAt, nil
}

func (s Session) NewRefreshTokenExpiry() (*time.Time, error) {
	refreshLifespan, err := strconv.Atoi(os.Getenv("REFRESH_LIFESPAN_MINUTES"))
	if err != nil {
		return nil, err
	}

	refreshExpiresAt := time.Now().Add(time.Minute * time.Duration(refreshLifespan))
	return &refreshExpiresAt, nil
}
