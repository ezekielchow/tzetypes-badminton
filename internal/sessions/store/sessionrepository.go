package sessions

import (
	"common/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, tx *pgx.Tx, userID string, sessionTokenExpiresAt time.Time, refreshTokenExpiresAt time.Time) (models.Session, error)
	FindSessionWithSessionID(ctx context.Context, tx *pgx.Tx, sessionID string) (models.Session, error)
	FindSessionToRefreshAccessToken(ctx context.Context, tx *pgx.Tx, refreshToken string) (models.Session, error)
	UpdateSessionWithRefreshToken(ctx context.Context, tx *pgx.Tx, refreshToken string, sessionTokenExpiresAt time.Time) (models.Session, error)
	DeleteSession(ctx context.Context, tx *pgx.Tx, sessionID string) error
}
