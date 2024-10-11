package player

import (
	"common/oapiprivate"
	"context"
	playerstore "players/store"
	userstore "users/store"
)

type PlayerServiceInterface interface {
	AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject) (oapiprivate.AddPlayerResponseObject, error)
	ListPlayers(ctx context.Context, input oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error)
}

type PlayerService struct {
	PlayerStore playerstore.PlayerRepository
	UserStore   userstore.UserRepository
}
