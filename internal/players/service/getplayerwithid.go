package players

import (
	"common/oapiprivate"
	"context"
	"database/sql"
	"errors"
	"strings"
)

const (
	GetPlayerWithIdNotFoundError = "unable to find player with given ID"
)

func (ps PlayerService) GetPlayerWithId(ctx context.Context, input oapiprivate.GetPlayerWithIdRequestObject) (oapiprivate.GetPlayerWithIdResponseObject, error) {

	player, err := ps.PlayerStore.GetPlayerWithId(ctx, nil, input.Id)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	if player.ID == "" {
		return nil, errors.New(GetPlayerWithIdNotFoundError)
	}

	return oapiprivate.GetPlayerWithId200JSONResponse{
		Id:        player.ID,
		Name:      player.Name,
		UserId:    player.UserID,
		CreatedAt: player.CreatedAt.String(),
		UpdatedAt: player.UpdatedAt.String(),
	}, nil
}
