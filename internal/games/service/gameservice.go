package games

import (
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	games "games/store"
	clubs "tzetypes-badminton/clubs/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GameServiceInterface interface {
	StartGame(ctx context.Context, input oapiprivate.StartGameRequestObject, user models.User) (oapiprivate.StartGameResponseObject, error)
	AddGameSteps(ctx context.Context, input oapiprivate.AddGameStepsRequestObject) (oapiprivate.AddGameStepsResponseObject, error)
	DeleteGameSteps(ctx context.Context, input oapiprivate.DeleteGameStepsRequestObject) (oapiprivate.DeleteGameStepsResponseObject, error)
	EndGame(ctx context.Context, input oapiprivate.EndGameRequestObject) (oapiprivate.EndGameResponseObject, error)
	GetGame(ctx context.Context, input oapiprivate.GetGameRequestObject) (oapiprivate.GetGameResponseObject, error)
	CreateOrUpdateGameHistory(ctx context.Context, input oapiprivate.CreateOrUpdateGameHistoryRequestObject, user models.User) (oapiprivate.CreateOrUpdateGameHistoryResponseObject, error)
	GetGameHistory(ctx context.Context, input oapiprivate.GetGameHistoryRequestObject, user models.User) (oapiprivate.GetGameHistoryResponseObject, error)
	GenerateRecentStatistics(ctx context.Context, input oapipublic.GenerateRecentStatisticsRequestObject) (oapipublic.GenerateRecentStatisticsResponseObject, error)
	GetRecentStatistics(ctx context.Context, input oapiprivate.GetRecentStatisticsRequestObject, user models.User) (oapiprivate.GetRecentStatisticsResponseObject, error)
	EndAbandonedGames(ctx context.Context, input oapipublic.EndAbandonedGamesRequestObject) (oapipublic.EndAbandonedGamesResponseObject, error)
	ListActiveGames(ctx context.Context, input oapiprivate.ListActiveGamesRequestObject, user models.User) (oapiprivate.ListActiveGamesResponseObject, error)
	GetGameStatistics(ctx context.Context, input oapipublic.GetGameStatisticsRequestObject) (oapipublic.GetGameStatisticsResponseObject, error)
	UpdateInstagramFeed(ctx context.Context, input oapipublic.UpdateInstagramFeedRequestObject) (oapipublic.UpdateInstagramFeedResponseObject, error)
	GetInstagramFeed(ctx context.Context, input oapipublic.GetInstagramFeedRequestObject) (oapipublic.GetInstagramFeedResponseObject, error)
}

type GameService struct {
	ClubStore clubs.ClubRepository
	GameStore games.GameRepository
	PgxPool   *pgxpool.Pool
}
