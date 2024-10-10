package sessionstore

import (
	"common/models"
	"context"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, userID string, sessionTokenExpiresAt time.Time, refreshTokenExpiresAt time.Time) (models.Session, error)
	FindSessionWithSessionID(ctx context.Context, sessionID string) (models.Session, error)
	FindSessionToRefreshAccessToken(ctx context.Context, refreshToken string) (models.Session, error)
	UpdateSessionWithRefreshToken(ctx context.Context, refreshToken string, sessionTokenExpiresAt time.Time) (models.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
