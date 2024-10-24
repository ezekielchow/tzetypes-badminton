package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"time"
)

func (gs GameService) StartGame(ctx context.Context, input oapiprivate.StartGameRequestObject, user models.User) (oapiprivate.StartGameResponseObject, error) {

	tx, err := gs.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	club, err := gs.ClubStore.GetClubGivenOwnerId(ctx, &tx, user.ID)
	if err != nil {
		return nil, err
	}

	leftOddPlayerName := ""
	if input.Body.LeftOddPlayerName != nil {
		leftOddPlayerName = *input.Body.LeftOddPlayerName
	}

	rightOddPlayerName := ""
	if input.Body.RightOddPlayerName != nil {
		rightOddPlayerName = *input.Body.RightOddPlayerName
	}

	game, err := gs.GameStore.CreateGame(ctx, &tx, models.Game{
		ClubID:              club.ID,
		LeftOddPlayerName:   &leftOddPlayerName,
		LeftEvenPlayerName:  input.Body.LeftEvenPlayerName,
		RightOddPlayerName:  &rightOddPlayerName,
		RightEvenPlayerName: input.Body.RightEvenPlayerName,
		GameType:            string(input.Body.GameType),
		ServingSide:         string(input.Body.ServingSide),
	})
	if err != nil {
		return nil, err
	}

	gameStep, err := gs.GameStore.CreateGameStep(ctx, &tx, models.GameStep{
		GameID:              game.ID,
		TeamLeftScore:       0,
		TeamRightScore:      0,
		ScoreAt:             time.Now(),
		StepNum:             1,
		CurrentServer:       string(input.Body.ServingSide),
		LeftOddPlayerName:   &leftOddPlayerName,
		LeftEvenPlayerName:  input.Body.LeftEvenPlayerName,
		RightOddPlayerName:  &rightOddPlayerName,
		RightEvenPlayerName: input.Body.RightEvenPlayerName,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.StartGame201JSONResponse{
		Game: oapiprivate.Game{
			ClubId:              game.ClubID,
			CreatedAt:           game.CreatedAt.String(),
			GameType:            game.GameType,
			Id:                  game.ID,
			LeftEvenPlayerName:  game.LeftEvenPlayerName,
			LeftOddPlayerName:   *game.LeftOddPlayerName,
			RightEvenPlayerName: game.RightEvenPlayerName,
			RightOddPlayerName:  *game.RightOddPlayerName,
			ServingSide:         game.ServingSide,
			UpdatedAt:           game.UpdatedAt.String(),
		},
		Steps: []oapiprivate.GameStep{
			{
				CreatedAt:           gameStep.CreatedAt.String(),
				GameId:              gameStep.GameID,
				Id:                  gameStep.ID,
				ScoreAt:             gameStep.ScoreAt.String(),
				StepNum:             gameStep.StepNum,
				TeamLeftScore:       gameStep.TeamLeftScore,
				TeamRightScore:      gameStep.TeamRightScore,
				CurrentServer:       gameStep.CurrentServer,
				LeftEvenPlayerName:  game.LeftEvenPlayerName,
				LeftOddPlayerName:   *game.LeftOddPlayerName,
				RightEvenPlayerName: game.RightEvenPlayerName,
				RightOddPlayerName:  *game.RightOddPlayerName,
				UpdatedAt:           gameStep.UpdatedAt.String(),
			},
		},
	}, nil
}
