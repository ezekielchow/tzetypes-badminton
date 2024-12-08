package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
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

		created, err := gs.GameStore.CreateGameStep(ctx, &tx, models.GameStep{
			GameID:              input.GameId,
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
			IsPaused:            step.IsPaused,
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
			IsPaused:            created.IsPaused,
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
