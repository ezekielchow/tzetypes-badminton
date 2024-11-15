package models

import (
	"common/oapiprivate"
	"common/oapipublic"
	"common/utils"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type TeamSide string

const (
	TeamSideLeft  TeamSide = "team_side_left"
	TeamSideRight TeamSide = "team_side_right"
)

type GameStep struct {
	ID                  string
	GameID              string
	TeamLeftScore       int
	TeamRightScore      int
	ScoreAt             time.Time
	StepNum             int
	CurrentServer       string
	LeftOddPlayerName   *string
	LeftEvenPlayerName  string
	RightOddPlayerName  *string
	RightEvenPlayerName string
	SyncId              string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

func (gs *GameStep) PostgresToModel(fromDb database.GameStep) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	gameID, err := uuid.FromBytes(fromDb.GameID.Bytes[:])
	if err != nil {
		return err
	}

	gs.ID = id.String()
	gs.GameID = gameID.String()
	gs.TeamLeftScore = int(fromDb.TeamLeftScore)
	gs.TeamRightScore = int(fromDb.TeamRightScore)
	gs.ScoreAt = fromDb.ScoreAt.Time
	gs.StepNum = int(fromDb.StepNum)
	gs.CreatedAt = fromDb.CreatedAt.Time
	gs.UpdatedAt = &fromDb.UpdatedAt.Time
	gs.CurrentServer = fromDb.CurrentServer
	gs.LeftEvenPlayerName = fromDb.LeftEvenPlayerName
	gs.LeftOddPlayerName = fromDb.LeftOddPlayerName
	gs.RightEvenPlayerName = fromDb.RightEvenPlayerName
	gs.RightOddPlayerName = fromDb.RightOddPlayerName
	gs.SyncId = fromDb.SyncID

	return nil
}

func (gs *GameStep) ModelToAPI() oapiprivate.GameStep {
	return oapiprivate.GameStep{
		CreatedAt:           gs.CreatedAt.String(),
		GameId:              gs.GameID,
		Id:                  gs.ID,
		ScoreAt:             gs.ScoreAt.String(),
		StepNum:             gs.StepNum,
		TeamLeftScore:       gs.TeamLeftScore,
		TeamRightScore:      gs.TeamRightScore,
		CurrentServer:       gs.CurrentServer,
		LeftEvenPlayerName:  gs.LeftEvenPlayerName,
		LeftOddPlayerName:   *gs.LeftOddPlayerName,
		RightEvenPlayerName: gs.RightEvenPlayerName,
		RightOddPlayerName:  *gs.RightOddPlayerName,
		UpdatedAt:           gs.UpdatedAt.String(),
		SyncId:              &gs.SyncId,
	}
}

func GameStepsToAPI(gameSteps []GameStep) []oapipublic.GameStep {
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

	return apiSteps
}

func GameStepFactory(count int, args map[string]interface{}) []GameStep {
	gameSteps := []GameStep{}

	gameID, ok := args["GameID"]
	if !ok {
		gameID = uuid.NewString()
	}

	teamLeftScore, ok := args["TeamLeftScore"]
	if !ok {
		teamLeftScore = 0
	}

	teamRightScore, ok := args["TeamRightScore"]
	if !ok {
		teamRightScore = 0
	}

	scoreAt, ok := args["ScoreAt"]
	if !ok {
		scoreAt = time.Now()
	}

	stepNum, ok := args["StepNum"]
	if !ok {
		stepNum = 0
	}

	leftOddPlayerName := utils.NewString(10)
	rightOddPlayerName := utils.NewString(10)

	for i := 0; i < count; i++ {
		gameSteps = append(gameSteps, GameStep{
			GameID:              gameID.(string),
			TeamLeftScore:       teamLeftScore.(int),
			TeamRightScore:      teamRightScore.(int),
			ScoreAt:             scoreAt.(time.Time),
			StepNum:             stepNum.(int),
			CurrentServer:       string(oapiprivate.LeftEven),
			LeftOddPlayerName:   &leftOddPlayerName,
			LeftEvenPlayerName:  utils.NewString(10),
			RightOddPlayerName:  &rightOddPlayerName,
			RightEvenPlayerName: utils.NewString(10),
			SyncId:              utils.NewString(10),
			CreatedAt:           time.Now(),
		})
	}

	return gameSteps
}
