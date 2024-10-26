package games

import (
	"common/oapiprivate"
	"context"
)

func (gs GameService) EndGame(ctx context.Context, input oapiprivate.EndGameRequestObject) (oapiprivate.EndGameResponseObject, error) {
	err := gs.GameStore.EndGame(ctx, nil, input.GameId, *input.Body.IsEnded)
	if err != nil {
		return nil, err
	}

	return oapiprivate.EndGame200Response{}, nil
}
