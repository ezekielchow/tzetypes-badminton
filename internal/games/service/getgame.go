package games

import (
	"common/oapipublic"
	"context"
)

func (gs GameService) GetGame(ctx context.Context, input oapipublic.GetGameRequestObject) (oapipublic.GetGameResponseObject, error) {

	game, err := gs.GameStore.GetGame(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	gameSteps, err := gs.GameStore.GetGameSteps(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	apiSteps := []oapipublic.GameStep{}
	for _, step := range gameSteps {
		apiSteps = append(apiSteps, oapipublic.GameStep{
			CreatedAt:           step.CreatedAt.String(),
			GameId:              step.GameID,
			Id:                  step.ID,
			ScoreAt:             step.ScoreAt.String(),
			StepNum:             step.StepNum,
			TeamLeftScore:       step.TeamLeftScore,
			TeamRightScore:      step.TeamRightScore,
			CurrentServer:       step.CurrentServer,
			LeftEvenPlayerName:  step.LeftEvenPlayerName,
			LeftOddPlayerName:   *step.LeftOddPlayerName,
			RightEvenPlayerName: step.RightEvenPlayerName,
			RightOddPlayerName:  *step.RightOddPlayerName,
			UpdatedAt:           step.UpdatedAt.String(),
			SyncId:              &step.SyncId,
		})
	}

	return oapipublic.GetGame200JSONResponse{
		Game: oapipublic.Game{
			ClubId:              game.ClubID,
			CreatedAt:           game.CreatedAt.String(),
			GameType:            game.GameType,
			Id:                  game.ID,
			LeftEvenPlayerName:  game.LeftEvenPlayerName,
			LeftOddPlayerName:   *game.LeftOddPlayerName,
			RightEvenPlayerName: game.RightEvenPlayerName,
			RightOddPlayerName:  *game.RightOddPlayerName,
			ServingSide:         game.ServingSide,
			IsEnded:             game.IsEnded,
			UpdatedAt:           game.UpdatedAt.String(),
		},
		Steps: apiSteps,
	}, nil
}
