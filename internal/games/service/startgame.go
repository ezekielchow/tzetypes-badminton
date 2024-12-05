package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	PlayersNameEmpty = "please fill in players name"
)

func returnStartGameError(err error) oapiprivate.StartGamedefaultJSONResponse {
	return oapiprivate.StartGamedefaultJSONResponse{
		Body:       oapiprivate.Error{Message: err.Error()},
		StatusCode: http.StatusBadRequest,
	}
}

func validateStartGame(input oapiprivate.StartGameRequestObject) error {
	if input.Body.GameType == oapiprivate.Singles && (len(input.Body.LeftEvenPlayerName) < 1 || len(input.Body.RightEvenPlayerName) < 1) {
		return errors.New(PlayersNameEmpty)
	}

	if input.Body.GameType == oapiprivate.Doubles && (len(input.Body.LeftEvenPlayerName) < 1 || len(*input.Body.LeftOddPlayerName) < 1 || len(input.Body.RightEvenPlayerName) < 1 || len(*input.Body.RightOddPlayerName) < 1) {
		return errors.New(PlayersNameEmpty)
	}

	return nil
}

func (gs GameService) StartGame(ctx context.Context, input oapiprivate.StartGameRequestObject, user models.User) (oapiprivate.StartGameResponseObject, error) {

	err := validateStartGame(input)
	if err != nil {
		return returnStartGameError(err), nil
	}

	tx, err := gs.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	club, err := gs.ClubStore.GetClubGivenOwnerID(ctx, &tx, user.ID)
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

	gameStartedAt := time.Now()

	game, err := gs.GameStore.CreateGame(ctx, &tx, models.Game{
		ClubID:              club.ID,
		LeftOddPlayerName:   &leftOddPlayerName,
		LeftEvenPlayerName:  input.Body.LeftEvenPlayerName,
		RightOddPlayerName:  &rightOddPlayerName,
		RightEvenPlayerName: input.Body.RightEvenPlayerName,
		GameType:            string(input.Body.GameType),
		ServingSide:         string(input.Body.ServingSide),
		CreatedAt:           gameStartedAt,
	})
	if err != nil {
		return nil, err
	}

	gameStep, err := gs.GameStore.CreateGameStep(ctx, &tx, models.GameStep{
		GameID:              game.ID,
		TeamLeftScore:       0,
		TeamRightScore:      0,
		ScoreAt:             gameStartedAt,
		StepNum:             1,
		CurrentServer:       string(input.Body.ServingSide),
		LeftOddPlayerName:   &leftOddPlayerName,
		LeftEvenPlayerName:  input.Body.LeftEvenPlayerName,
		RightOddPlayerName:  &rightOddPlayerName,
		RightEvenPlayerName: input.Body.RightEvenPlayerName,
		SyncId:              uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.StartGame201JSONResponse{
		StartGame201ResponseSchemaJSONResponse: oapiprivate.StartGame201ResponseSchemaJSONResponse{
			Game: game.ModelToAPI(),
			Steps: []oapiprivate.GameStep{
				gameStep.ModelToAPI(),
			},
		},
	}, nil
}
