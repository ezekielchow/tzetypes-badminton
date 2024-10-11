package player

import (
	"common/oapiprivate"
	"context"
	playerstore "players/store"
	userstore "users/store"
)

type PlayerServiceInterface interface {
	AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject) (oapiprivate.AddPlayerResponseObject, error)
}

type PlayerService struct {
	PlayerStore playerstore.PlayerRepository
	UserStore   userstore.UserRepository
}
