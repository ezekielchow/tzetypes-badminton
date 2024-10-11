package player

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (ps PlayerService) AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject) (oapiprivate.AddPlayerResponseObject, error) {

	user, err := ps.UserStore.Signup(ctx, "", "")
	if err != nil {
		return nil, err
	}

	_, err = ps.PlayerStore.CreatePlayer(ctx, models.Player{
		UserID: user.ID,
		Name:   input.Body.Name,
	}, "")

	if err != nil {
		return nil, err
	}

	return oapiprivate.AddPlayer201Response{}, nil
}
