package database

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // package does so
	_ "github.com/jackc/pgx/v5"                                // package does so
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct{}

func (d Database) RunMigrations(dbURI string, migrationFiles string) error {

	m, err := migrate.New(
		migrationFiles,
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
