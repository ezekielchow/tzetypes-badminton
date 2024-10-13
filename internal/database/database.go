package database

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct{}

func (d Database) RunMigrations(dbURI string) error {

	m, err := migrate.New(
		"file://database/migrations",
		dbURI)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d Database) Open(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection is established
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
