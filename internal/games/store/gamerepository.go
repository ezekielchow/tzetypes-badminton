package games

import (
	"common/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type GameRepository interface {
	CreateGame(ctx context.Context, tx *pgx.Tx, toCreate models.Game) (models.Game, error)
	CreateGameStep(ctx context.Context, tx *pgx.Tx, toCreate models.GameStep) (models.GameStep, error)
	DeleteGameStep(ctx context.Context, tx *pgx.Tx, id string) error
	EndGame(ctx context.Context, tx *pgx.Tx, id string, isEnded bool) error
	GetGameSteps(ctx context.Context, tx *pgx.Tx, gameID string) ([]models.GameStep, error)
	GetGame(ctx context.Context, tx *pgx.Tx, id string) (models.Game, error)
	CreateStatistic(ctx context.Context, tx *pgx.Tx, gameID string, toCreate models.GameStatistic) (models.GameStatistic, error)
	GetStatisticsWithGameId(ctx context.Context, tx *pgx.Tx, gameID string) (models.GameStatistic, error)
	CreateOrUpdateGameHistory(ctx context.Context, tx *pgx.Tx, toCreate models.GameHistory) (models.GameHistory, error)
	GetGameHistoryGivenUserIdAndGameId(ctx context.Context, tx *pgx.Tx, userID string, gameID string) (models.GameHistory, error)
}
