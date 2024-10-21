package players

import (
	"common/oapiprivate"
	"context"
	"errors"
)

const (
	GetPlayerWithIdNotFoundError = "unable to find player with given ID"
)

func (ps PlayerService) GetPlayerWithId(ctx context.Context, input oapiprivate.GetPlayersIdRequestObject) (oapiprivate.GetPlayersIdResponseObject, error) {

	player, err := ps.PlayerStore.GetPlayerWithId(ctx, nil, input.Id)
	if err != nil {
		return nil, err
	}

	if player.ID == "" {
		return nil, errors.New(GetPlayerWithIdNotFoundError)
	}

	return oapiprivate.GetPlayersId200JSONResponse{
		Id:        player.ID,
		Name:      player.Name,
		UserId:    player.UserID,
		CreatedAt: player.CreatedAt.String(),
		UpdatedAt: player.UpdatedAt.String(),
	}, nil
}
