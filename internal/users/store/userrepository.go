package users

import (
	"common/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *pgx.Tx, firebaseUID string, email string, accountTier string) (models.User, error)
	FindUserWithEmail(ctx context.Context, tx *pgx.Tx, email string) (models.User, error)
	FindUserWithID(ctx context.Context, tx *pgx.Tx, id string) (models.User, error)
}
