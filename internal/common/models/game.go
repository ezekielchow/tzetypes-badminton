package models

import (
	"common/oapiprivate"
	"common/oapipublic"
	"common/utils"
	"fmt"
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

func (g *Game) ModelToAPI() oapiprivate.Game {
	return oapiprivate.Game{
		ClubId:              g.ClubID,
		CreatedAt:           g.CreatedAt.String(),
		GameType:            g.GameType,
		Id:                  g.ID,
		LeftEvenPlayerName:  g.LeftEvenPlayerName,
		LeftOddPlayerName:   *g.LeftOddPlayerName,
		RightEvenPlayerName: g.RightEvenPlayerName,
		RightOddPlayerName:  *g.RightOddPlayerName,
		ServingSide:         g.ServingSide,
		IsEnded:             g.IsEnded,
		UpdatedAt:           g.UpdatedAt.String(),
	}
}

func GetGameLengthFormatted(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60

	return fmt.Sprintf("%02d hours %02d minutes", hours, minutes)
}

func (g *Game) ModelToPublicAPI() oapipublic.Game {
	return oapipublic.Game{
		ClubId:              g.ClubID,
		CreatedAt:           g.CreatedAt.String(),
		GameType:            g.GameType,
		Id:                  g.ID,
		LeftEvenPlayerName:  g.LeftEvenPlayerName,
		LeftOddPlayerName:   *g.LeftOddPlayerName,
		RightEvenPlayerName: g.RightEvenPlayerName,
		RightOddPlayerName:  *g.RightOddPlayerName,
		ServingSide:         g.ServingSide,
		IsEnded:             g.IsEnded,
		UpdatedAt:           g.UpdatedAt.String(),
	}
}

func GameFactory(count int, args map[string]interface{}) []Game {
	games := []Game{}

	clubID, ok := args["ClubID"]
	if !ok {
		clubID = utils.NewString(10)
	}

	gameType, ok := args["GameType"]
	if !ok {
		gameType = string(oapiprivate.Singles)
	}

	rightOddName := ""
	leftOddName := ""
	if gameType == oapiprivate.Doubles {
		rightOddName = utils.NewString(10)
		leftOddName = utils.NewString(10)
	}

	servingSide, ok := args["ServingSide"]
	if !ok {
		servingSide = string(oapiprivate.RightEven)
	}

	for i := 0; i < count; i++ {
		games = append(games, Game{
			ClubID:              clubID.(string),
			LeftOddPlayerName:   &leftOddName,
			LeftEvenPlayerName:  utils.NewString(10),
			RightOddPlayerName:  &rightOddName,
			RightEvenPlayerName: utils.NewString(10),
			GameType:            gameType.(string),
			ServingSide:         servingSide.(string),
			IsEnded:             false,
			CreatedAt:           time.Now(),
		})
	}

	return games
}
