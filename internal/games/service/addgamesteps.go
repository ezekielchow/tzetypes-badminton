package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"log"
	"time"
)

func (gs GameService) AddGameSteps(ctx context.Context, input oapiprivate.AddGameStepsRequestObject) (oapiprivate.AddGameStepsResponseObject, error) {

	tx, err := gs.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	stepsRes := []oapiprivate.GameStep{}
	for _, step := range input.Body.Steps {
		scoreTime, err := time.Parse(time.RFC3339, step.ScoreAt)
		if err != nil {
			return nil, err
		}

		log.Println("AAAAAAA", *step.SyncId)

		created, err := gs.GameStore.CreateGameStep(ctx, &tx, models.GameStep{
			GameID:              input.Id,
			TeamLeftScore:       step.TeamLeftScore,
			TeamRightScore:      step.TeamRightScore,
			ScoreAt:             scoreTime,
			StepNum:             step.StepNum,
			CurrentServer:       step.CurrentServer,
			LeftOddPlayerName:   &step.LeftOddPlayerName,
			LeftEvenPlayerName:  step.LeftEvenPlayerName,
			RightOddPlayerName:  &step.RightOddPlayerName,
			RightEvenPlayerName: step.RightEvenPlayerName,
			SyncId:              *step.SyncId,
		})
		if err != nil {
			return nil, err
		}

		stepsRes = append(stepsRes, oapiprivate.GameStep{
			CurrentServer:       created.CurrentServer,
			GameId:              created.GameID,
			Id:                  created.ID,
			LeftEvenPlayerName:  created.LeftEvenPlayerName,
			LeftOddPlayerName:   *created.LeftOddPlayerName,
			RightEvenPlayerName: created.RightEvenPlayerName,
			RightOddPlayerName:  *created.RightOddPlayerName,
			ScoreAt:             created.ScoreAt.String(),
			StepNum:             created.StepNum,
			SyncId:              &created.SyncId,
			TeamLeftScore:       created.TeamLeftScore,
			TeamRightScore:      created.TeamRightScore,
		})
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.AddGameSteps201JSONResponse{
		GameSteps: stepsRes,
	}, nil
}
