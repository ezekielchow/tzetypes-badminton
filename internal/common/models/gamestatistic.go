package models

import (
	"common/oapipublic"
	"fmt"
	"math"
	"strconv"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type GameStatistic struct {
	ID                            string
	GameID                        string
	TotalGameTimeSeconds          int
	RightConsecutivePointsSeconds int
	LeftConsecutivePointsSeconds  int
	LongestPointSeconds           int
	ShortestPointSeconds          int
	AverageTimePerPointSeconds    int
	CreatedAt                     time.Time
	UpdatedAt                     *time.Time
}

func (gs *GameStatistic) PostgresToModel(fromDb database.GameStatistic) error {
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
	gs.TotalGameTimeSeconds = int(*fromDb.TotalGameTimeSeconds)
	gs.RightConsecutivePointsSeconds = int(*fromDb.RightConsecutivePointsSeconds)
	gs.LeftConsecutivePointsSeconds = int(*fromDb.LeftConsecutivePointsSeconds)
	gs.LongestPointSeconds = int(*fromDb.LongestPointSeconds)
	gs.ShortestPointSeconds = int(*fromDb.ShortestPointSeconds)
	gs.AverageTimePerPointSeconds = int(*fromDb.AverageTimePerPointSeconds)
	gs.CreatedAt = fromDb.CreatedAt.Time
	gs.UpdatedAt = &fromDb.UpdatedAt.Time
	return nil
}

func (gs *GameStatistic) ModelToAPI() oapipublic.GameStatistic {
	return oapipublic.GameStatistic{
		AveragePerPoint:        fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(gs.AverageTimePerPointSeconds)/60), gs.AverageTimePerPointSeconds%60),
		LeftConsecutivePoints:  strconv.Itoa(gs.LeftConsecutivePointsSeconds),
		LongestPoint:           fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(gs.LongestPointSeconds)/60), gs.LongestPointSeconds%60),
		RightConsecutivePoints: strconv.Itoa(gs.RightConsecutivePointsSeconds),
		ShortestPoint:          fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(gs.ShortestPointSeconds)/60), gs.ShortestPointSeconds%60),
		TotalGameTime:          fmt.Sprintf("%02.fh %02.dm", math.Floor(float64(gs.TotalGameTimeSeconds)/60/60), (gs.TotalGameTimeSeconds/60)%60),
	}
}
