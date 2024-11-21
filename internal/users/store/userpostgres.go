package users

import (
	"common/models"
	"common/utils"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
)

type UserPostgres struct {
	Queries *database.Queries
}

func (up UserPostgres) CreateUser(ctx context.Context, tx *pgx.Tx, email string, passwordHash string, userType string) (models.User, error) {
	queries := up.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	created, err := queries.CreateUser(ctx, database.CreateUserParams{
		Email:        email,
		PasswordHash: &passwordHash,
		UserType:     userType,
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

	pgID, err := utils.StringToPgId(id)
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
