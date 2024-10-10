package userstore

import (
	"common/models"
	"context"
)

type UserRepository interface {
	Signup(ctx context.Context, email string, passwordHash string) (models.User, error)
	FindUserWithEmail(ctx context.Context, email string) (models.User, error)
	FindUserWithID(ctx context.Context, id string) (models.User, error)
}
