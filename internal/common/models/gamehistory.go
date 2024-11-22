package models

import (
	"common/oapiprivate"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type GameHistory struct {
	ID             string
	UserID         string
	GameID         string
	PlayerPosition string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (gh *GameHistory) PostgresToModel(fromDb database.GameHistory) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	userID, err := uuid.FromBytes(fromDb.UserID.Bytes[:])
	if err != nil {
		return err
	}

	gameID, err := uuid.FromBytes(fromDb.GameID.Bytes[:])
	if err != nil {
		return err
	}

	gh.ID = id.String()
	gh.UserID = userID.String()
	gh.GameID = gameID.String()
	gh.PlayerPosition = fromDb.PlayerPosition
	gh.CreatedAt = fromDb.CreatedAt.Time
	gh.UpdatedAt = &fromDb.UpdatedAt.Time
	return nil
}

func (gh *GameHistory) ModelToAPI() oapiprivate.GameHistory {
	return oapiprivate.GameHistory{
		Id:             gh.ID,
		UserId:         gh.UserID,
		GameId:         gh.GameID,
		PlayerPosition: gh.PlayerPosition,
		CreatedAt:      gh.CreatedAt.String(),
		UpdatedAt:      gh.UpdatedAt.String(),
	}
}
