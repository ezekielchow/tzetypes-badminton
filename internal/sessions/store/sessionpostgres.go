package sessionstore

import (
	"common/models"
	"context"
	sessionstore "sessions/store/generated"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SessionPostgres struct {
	Queries *sessionstore.Queries
}

func (sp SessionPostgres) CreateSession(ctx context.Context, userID string, sessionTokenExpiry time.Time, refreshTokenExpiry time.Time) (models.Session, error) {
	pgID := pgtype.UUID{}
	err := pgID.Scan(userID)
	if err != nil {
		return models.Session{}, err
	}

	sessionTokenExpiryTimestamp := pgtype.Timestamp{}
	err = sessionTokenExpiryTimestamp.Scan(sessionTokenExpiry)
	if err != nil {
		return models.Session{}, err
	}

	refreshTokenExpiryTimestamp := pgtype.Timestamp{}
	err = refreshTokenExpiryTimestamp.Scan(refreshTokenExpiry)
	if err != nil {
		return models.Session{}, err
	}

	created, err := sp.Queries.CreateSession(ctx, sessionstore.CreateSessionParams{
		UserID:                pgID,
		SessionTokenExpiresAt: sessionTokenExpiryTimestamp,
		RefreshTokenExpiresAt: refreshTokenExpiryTimestamp,
	})

	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{}
	err = session.PostgresToModel(created)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (sp SessionPostgres) FindSessionWithSessionID(ctx context.Context, sessionID string) (models.Session, error) {
	pgID := pgtype.UUID{}
	err := pgID.Scan(sessionID)
	if err != nil {
		return models.Session{}, err
	}

	found, err := sp.Queries.FindSessionWithSessionID(ctx, pgID)
	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{}
	err = session.PostgresToModel(found)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (sp SessionPostgres) FindSessionToRefreshAccessToken(ctx context.Context, refreshToken string) (models.Session, error) {

	refreshTokenPG := pgtype.UUID{}
	err := refreshTokenPG.Scan(refreshToken)
	if err != nil {
		return models.Session{}, err
	}

	found, err := sp.Queries.FindSessionWithRefreshToken(ctx, refreshTokenPG)
	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{}
	err = session.PostgresToModel(found)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (sp SessionPostgres) UpdateSessionWithRefreshToken(ctx context.Context, refreshToken string, sessionTokenExpiresAt time.Time) (models.Session, error) {
	refreshTokenPG := pgtype.UUID{}
	err := refreshTokenPG.Scan(refreshToken)
	if err != nil {
		return models.Session{}, err
	}

	sessionTokenExpiry := pgtype.Timestamp{}
	err = sessionTokenExpiry.Scan(sessionTokenExpiresAt)
	if err != nil {
		return models.Session{}, err
	}

	updated, err := sp.Queries.UpdateSessionWithRefreshToken(ctx, sessionstore.UpdateSessionWithRefreshTokenParams{
		RefreshToken:          refreshTokenPG,
		SessionTokenExpiresAt: sessionTokenExpiry,
	})
	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{}
	err = session.PostgresToModel(updated)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (sp SessionPostgres) DeleteSession(ctx context.Context, sessionID string) error {
	pgID := pgtype.UUID{}
	err := pgID.Scan(sessionID)
	if err != nil {
		return err
	}

	err = sp.Queries.DeleteSession(ctx, pgID)
	if err != nil {
		return err
	}

	return nil
}
