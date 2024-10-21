package players

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
	UpdatePlayer(ctx context.Context, input oapiprivate.UpdatePlayerWithIdRequestObject) (oapiprivate.UpdatePlayerWithIdResponseObject, error)
	GetPlayerWithId(ctx context.Context, input oapiprivate.GetPlayerWithIdRequestObject) (oapiprivate.GetPlayerWithIdResponseObject, error)
}

type PlayerService struct {
	PlayerStore playerstore.PlayerRepository
	UserStore   userstore.UserRepository
	ClubStore   clubs.ClubRepository
	PgxPool     *pgxpool.Pool
}
