package models

import (
	"common/oapiprivate"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type GameRecentStatistic struct {
	ID                             string
	UserID                         string
	GameCount                      int
	Wins                           int
	Losses                         int
	TotalPoints                    int
	PointsWon                      int
	AverageTimePerPointSeconds     int
	AverageTimePerPointWonSeconds  int
	AverageTimePerPointLostSeconds int
	LongestRallySeconds            int
	LongestRallyIsWon              int
	ShortestRallySeconds           int
	ShortestRallyIsWon             int
	NeedsRegenerating              int
	CreatedAt                      time.Time
	UpdatedAt                      *time.Time
}

func (grs *GameRecentStatistic) PostgresToModel(fromDb database.GameRecentStatistic) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	userID, err := uuid.FromBytes(fromDb.UserID.Bytes[:])
	if err != nil {
		return err
	}

	grs.ID = id.String()
	grs.UserID = userID.String()
	grs.GameCount = int(*fromDb.GameCount)
	grs.Wins = int(*fromDb.Wins)
	grs.Losses = int(*fromDb.Losses)
	grs.TotalPoints = int(*fromDb.TotalPoints)
	grs.PointsWon = int(*fromDb.PointsWon)
	grs.AverageTimePerPointSeconds = int(*fromDb.AverageTimePerPointSeconds)
	grs.AverageTimePerPointWonSeconds = int(*fromDb.AverageTimePerPointWonSeconds)
	grs.AverageTimePerPointLostSeconds = int(*fromDb.AverageTimePerPointLostSeconds)
	grs.LongestRallySeconds = int(*fromDb.LongestRallySeconds)
	grs.LongestRallyIsWon = int(*fromDb.LongestRallyIsWon)
	grs.ShortestRallySeconds = int(*fromDb.ShortestRallySeconds)
	grs.ShortestRallyIsWon = int(*fromDb.ShortestRallyIsWon)
	grs.NeedsRegenerating = int(*fromDb.NeedsRegenerating)
	grs.CreatedAt = fromDb.CreatedAt.Time
	grs.UpdatedAt = &fromDb.UpdatedAt.Time
	return nil
}

func (grs *GameRecentStatistic) ModelToAPI() oapiprivate.GameRecentStatistic {
	return oapiprivate.GameRecentStatistic{
		AverageTimePerPointLostSeconds: grs.AverageTimePerPointLostSeconds,
		AverageTimePerPointSeconds:     grs.AverageTimePerPointSeconds,
		AverageTimePerPointWonSeconds:  grs.AverageTimePerPointWonSeconds,
		CreatedAt:                      grs.CreatedAt.String(),
		GameCount:                      grs.GameCount,
		Id:                             grs.ID,
		LongestRallyIsWon:              grs.LongestRallyIsWon,
		LongestRallySeconds:            grs.LongestRallySeconds,
		Losses:                         grs.Losses,
		PointsWon:                      grs.PointsWon,
		ShortestRallyIsWon:             grs.ShortestRallyIsWon,
		ShortestRallySeconds:           grs.ShortestRallySeconds,
		TotalPoints:                    grs.TotalPoints,
		UpdatedAt:                      grs.UpdatedAt.String(),
		UserId:                         grs.UserID,
		Wins:                           grs.Wins,
	}
}
