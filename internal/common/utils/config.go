package utils

import (
	"errors"
	"os"
)

var ENVS = []string{
	APP_PORT,
	CORS_ALLOWED_ORIGINS,
	POSTGRES_PASSWORD,
	POSTGRES_USER,
	POSTGRES_DB,
	POSTGRES_HOST,
	SESSION_LIFESPAN_MINUTES,
	REFRESH_LIFESPAN_MINUTES,
	IS_HTTPS,
}

var (
	APP_PORT                 = "APP_PORT"
	CORS_ALLOWED_ORIGINS     = "CORS_ALLOWED_ORIGINS"
	POSTGRES_PASSWORD        = "POSTGRES_PASSWORD"
	POSTGRES_USER            = "POSTGRES_USER"
	POSTGRES_DB              = "POSTGRES_DB"
	POSTGRES_HOST            = "POSTGRES_HOST"
	SESSION_LIFESPAN_MINUTES = "SESSION_LIFESPAN_MINUTES"
	REFRESH_LIFESPAN_MINUTES = "REFRESH_LIFESPAN_MINUTES"
	IS_HTTPS                 = "IS_HTTPS"
)

type Config struct{}

func (c Config) LoadEnv() error {
	for _, key := range ENVS {
		if os.Getenv(key) == "" {
			return errors.New("Env is missing: " + key)
		}
	}

	return nil
}

func (c Config) Get(key string) string {
	return os.Getenv(key)
}
