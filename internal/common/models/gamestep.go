package models

import (
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
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
