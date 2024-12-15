package games

import (
	"common/models"
	"context"
	database "tzetypes-badminton/database/generated"

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
	CreateOrUpdateGameRecentStatistic(ctx context.Context, tx *pgx.Tx, toCreate models.GameRecentStatistic) (models.GameRecentStatistic, error)
	GetGameRecentStatisticWithUserId(ctx context.Context, tx *pgx.Tx, userID string) (models.GameRecentStatistic, error)
	GetGameRecentStatisticThatNeedsRegeneration(ctx context.Context, tx *pgx.Tx) ([]models.GameRecentStatistic, error)
	GetMostRecentGameHistories(ctx context.Context, tx *pgx.Tx, userID string) ([]models.GameHistory, error)
	GetGameStepsGivenGameIds(ctx context.Context, tx *pgx.Tx, gameIDs []string) ([]models.GameStep, error)
	GetAbandonedGames(ctx context.Context, tx *pgx.Tx) ([]string, error)
	EndGames(ctx context.Context, tx *pgx.Tx, ids []string) error
	GetActiveGames(ctx context.Context, tx *pgx.Tx, clubID string) ([]models.Game, error)
	GetPlayedGames(ctx context.Context, tx *pgx.Tx, userID string, sortIsGameWon string, sortGameCreatedAt string, offset int, limit int) ([]database.GetPlayedGamesRow, error)
	GetClubGames(ctx context.Context, tx *pgx.Tx, clubID string, sortGameCreatedAt string, offset int, limit int) ([]database.GetClubGamesRow, error)
	UpdateInstagramFeed(ctx context.Context, tx *pgx.Tx, media models.InstagramMedia) error
	GetInstagramFeedCount(ctx context.Context, tx *pgx.Tx) (int64, error)
}
