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
}

type GameService struct {
	ClubStore clubs.ClubRepository
	GameStore games.GameRepository
	PgxPool   *pgxpool.Pool
}
