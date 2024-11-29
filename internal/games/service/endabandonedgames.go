package games

import (
	"common/oapipublic"
	"context"
)

func (gs GameService) EndAbandonedGames(ctx context.Context, input oapipublic.EndAbandonedGamesRequestObject) (oapipublic.EndAbandonedGamesResponseObject, error) {
	ids, err := gs.GameStore.GetAbandonedGames(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = gs.GameStore.EndGames(ctx, nil, ids)
	if err != nil {
		return nil, err
	}

	return oapipublic.EndAbandonedGames200Response{}, nil
}
