package userstore

import (
	"common/models"
	"context"
	users "users/store/generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserPostgres struct {
	Queries *users.Queries
}

func (up UserPostgres) Signup(ctx context.Context, email string, passwordHash string) (models.User, error) {
	created, err := up.Queries.CreateUser(ctx, users.CreateUserParams{
		Email:        email,
		PasswordHash: &passwordHash,
	})

	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = user.PostgresToModel(created)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (up UserPostgres) FindUserWithEmail(ctx context.Context, email string) (models.User, error) {
	res, err := up.Queries.FindUserWithEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	user.PostgresToModel(res)

	return user, nil
}

func (up UserPostgres) FindUserWithID(ctx context.Context, id string) (models.User, error) {
	pgID := pgtype.UUID{}
	err := pgID.Scan(id)
	if err != nil {
		return models.User{}, err
	}

	res, err := up.Queries.FindUserWithID(ctx, pgID)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	user.PostgresToModel(res)

	return user, nil
}
