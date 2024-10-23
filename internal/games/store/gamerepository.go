package games

import (
	"common/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type GameRepository interface {
	CreateGame(ctx context.Context, tx *pgx.Tx, toCreate models.Game) (models.Game, error)
	CreateGameStep(ctx context.Context, tx *pgx.Tx, toCreate models.GameStep) (models.GameStep, error)
}
