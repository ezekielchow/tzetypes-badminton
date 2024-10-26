package games

import (
	"common/models"
	"common/oapiprivate"
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
}

type GameService struct {
	ClubStore clubs.ClubRepository
	GameStore games.GameRepository
	PgxPool   *pgxpool.Pool
}
