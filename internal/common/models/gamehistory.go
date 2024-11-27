package models

import (
	"common/oapiprivate"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type GameHistory struct {
	ID                             string
	UserID                         string
	GameID                         string
	PlayerPosition                 string
	GameStartedAt                  time.Time
	GameWonBy                      string
	TotalPoints                    int
	PointsWon                      int
	PointsLost                     int
	AverageTimePerPointSeconds     int
	AverageTimePerPointWonSeconds  int
	AverageTimePerPointLostSeconds int
	LongestRallySeconds            int
	LongestRallyIsWon              int
	ShortestRallySeconds           int
	ShortestRallyIsWon             int
	IsGameWon                      int
	CreatedAt                      time.Time
	UpdatedAt                      *time.Time
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
	gh.GameStartedAt = fromDb.GameStartedAt.Time
	gh.GameWonBy = fromDb.GameWonBy
	gh.TotalPoints = int(fromDb.TotalPoints)
	gh.PointsWon = int(fromDb.PointsWon)
	gh.PointsLost = int(fromDb.PointsLost)
	gh.AverageTimePerPointSeconds = int(fromDb.AverageTimePerPointSeconds)
	gh.AverageTimePerPointWonSeconds = int(fromDb.AverageTimePerPointWonSeconds)
	gh.AverageTimePerPointLostSeconds = int(fromDb.AverageTimePerPointLostSeconds)
	gh.LongestRallySeconds = int(fromDb.LongestRallySeconds)
	gh.LongestRallyIsWon = int(fromDb.LongestRallyIsWon)
	gh.ShortestRallySeconds = int(fromDb.ShortestRallySeconds)
	gh.ShortestRallyIsWon = int(fromDb.ShortestRallyIsWon)
	gh.IsGameWon = int(fromDb.IsGameWon)
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
