package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (gs GameService) GetGame(ctx context.Context, input oapiprivate.GetGameRequestObject) (oapiprivate.GetGameResponseObject, error) {

	game, err := gs.GameStore.GetGame(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	gameSteps, err := gs.GameStore.GetGameSteps(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	apiSteps := models.GameStepsToPrivateAPI(gameSteps)

	return oapiprivate.GetGame200JSONResponse{
		GetGame200ResponseSchemaJSONResponse: oapiprivate.GetGame200ResponseSchemaJSONResponse{
			Game:  game.ModelToAPI(),
			Steps: apiSteps,
		},
	}, nil
}
