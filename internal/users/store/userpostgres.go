package userstore

import (
	"common/models"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserPostgres struct {
	Queries *database.Queries
}

func (up UserPostgres) CreateUser(ctx context.Context, tx *pgx.Tx, email string, passwordHash string) (models.User, error) {
	queries := up.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	created, err := queries.CreateUser(ctx, database.CreateUserParams{
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

func (up UserPostgres) FindUserWithEmail(ctx context.Context, tx *pgx.Tx, email string) (models.User, error) {
	queries := up.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	res, err := queries.FindUserWithEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	user.PostgresToModel(res)

	return user, nil
}

func (up UserPostgres) FindUserWithID(ctx context.Context, tx *pgx.Tx, id string) (models.User, error) {
	queries := up.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID := pgtype.UUID{}
	err := pgID.Scan(id)
	if err != nil {
		return models.User{}, err
	}

	res, err := queries.FindUserWithID(ctx, pgID)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	user.PostgresToModel(res)

	return user, nil
}
