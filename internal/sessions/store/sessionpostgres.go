package sessions

import (
	"common/models"
	"context"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SessionPostgres struct {
	Queries *database.Queries
}

func (sp SessionPostgres) CreateSession(ctx context.Context, tx *pgx.Tx, userID string, sessionTokenExpiry time.Time, refreshTokenExpiry time.Time) (models.Session, error) {
	queries := sp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

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

	created, err := queries.CreateSession(ctx, database.CreateSessionParams{
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

func (sp SessionPostgres) FindSessionWithSessionID(ctx context.Context, tx *pgx.Tx, sessionID string) (models.Session, error) {
	queries := sp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID := pgtype.UUID{}
	err := pgID.Scan(sessionID)
	if err != nil {
		return models.Session{}, err
	}

	found, err := queries.FindSessionWithSessionID(ctx, pgID)
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

func (sp SessionPostgres) FindSessionToRefreshAccessToken(ctx context.Context, tx *pgx.Tx, refreshToken string) (models.Session, error) {
	queries := sp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	refreshTokenPG := pgtype.UUID{}
	err := refreshTokenPG.Scan(refreshToken)
	if err != nil {
		return models.Session{}, err
	}

	found, err := queries.FindSessionWithRefreshToken(ctx, refreshTokenPG)
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

func (sp SessionPostgres) UpdateSessionWithRefreshToken(ctx context.Context, tx *pgx.Tx, refreshToken string, sessionTokenExpiresAt time.Time) (models.Session, error) {
	queries := sp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

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

	updated, err := queries.UpdateSessionWithRefreshToken(ctx, database.UpdateSessionWithRefreshTokenParams{
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

func (sp SessionPostgres) DeleteSession(ctx context.Context, tx *pgx.Tx, sessionID string) error {
	queries := sp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID := pgtype.UUID{}
	err := pgID.Scan(sessionID)
	if err != nil {
		return err
	}

	err = queries.DeleteSession(ctx, pgID)
	if err != nil {
		return err
	}

	return nil
}
