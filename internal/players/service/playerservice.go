package player

import (
	"common/oapiprivate"
	"context"
	playerstore "players/store"
	clubs "tzetypes-badminton/clubs/store"
	userstore "users/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlayerServiceInterface interface {
	AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject, ownerID string) (oapiprivate.AddPlayerResponseObject, error)
	ListPlayers(ctx context.Context, input oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error)
}

type PlayerService struct {
	PlayerStore playerstore.PlayerRepository
	UserStore   userstore.UserRepository
	ClubStore   clubs.ClubRepository
	PgxPool     *pgxpool.Pool
}
