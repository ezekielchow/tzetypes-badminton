package models

import (
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type Game struct {
	ID                  string
	ClubID              string
	LeftOddPlayerName   *string
	LeftEvenPlayerName  string
	RightOddPlayerName  *string
	RightEvenPlayerName string
	GameType            string
	ServingSide         string
	IsEnded             bool
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

func (g *Game) PostgresToModel(fromDb database.Game) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	clubID, err := uuid.FromBytes(fromDb.ClubID.Bytes[:])
	if err != nil {
		return err
	}

	g.ID = id.String()
	g.ClubID = clubID.String()
	g.LeftEvenPlayerName = fromDb.LeftEvenPlayerName
	g.LeftOddPlayerName = fromDb.LeftOddPlayerName
	g.RightEvenPlayerName = fromDb.RightEvenPlayerName
	g.RightOddPlayerName = fromDb.RightOddPlayerName
	g.GameType = fromDb.GameType
	g.ServingSide = fromDb.ServingSide
	g.CreatedAt = fromDb.CreatedAt.Time
	g.UpdatedAt = &fromDb.UpdatedAt.Time
	g.IsEnded = fromDb.IsEnded

	return nil
}
