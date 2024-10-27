package games

import (
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

	apiSteps := []oapiprivate.GameStep{}
	for _, step := range gameSteps {
		apiSteps = append(apiSteps, step.ModelToAPI())
	}

	return oapiprivate.GetGame200JSONResponse{
		Game:  game.ModelToAPI(),
		Steps: apiSteps,
	}, nil
}
