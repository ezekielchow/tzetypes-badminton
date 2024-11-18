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
	ID                              string
	GameID                          string
	TotalGameTimeSeconds            int
	RightConsecutivePoints          int
	LeftConsecutivePoints           int
	LongestPointSeconds             int
	LongestPointTeam                string
	ShortestPointSeconds            int
	ShortestPointTeam               string
	AverageTimePerPointSeconds      int
	LeftAverageTimePerPointSeconds  int
	RightAverageTimePerPointSeconds int
	CreatedAt                       time.Time
	UpdatedAt                       *time.Time
}

type FormattedGameStatistic struct {
	TotalGameTimeSeconds            string
	RightConsecutivePoints          string
	LeftConsecutivePoints           string
	LongestPointSeconds             string
	LongestPointTeam                string
	ShortestPointSeconds            string
	ShortestPointTeam               string
	AverageTimePerPointSeconds      string
	LeftAverageTimePerPointSeconds  string
	RightAverageTimePerPointSeconds string
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
	gs.RightConsecutivePoints = int(*fromDb.RightConsecutivePoints)
	gs.LeftConsecutivePoints = int(*fromDb.LeftConsecutivePoints)
	gs.LongestPointSeconds = int(*fromDb.LongestPointSeconds)
	gs.LongestPointTeam = *fromDb.LongestPointTeam
	gs.ShortestPointSeconds = int(*fromDb.ShortestPointSeconds)
	gs.ShortestPointTeam = *fromDb.ShortestPointTeam
	gs.AverageTimePerPointSeconds = int(*fromDb.AverageTimePerPointSeconds)
	gs.RightAverageTimePerPointSeconds = int(*fromDb.RightAverageTimePerPointSeconds)
	gs.LeftAverageTimePerPointSeconds = int(*fromDb.LeftAverageTimePerPointSeconds)
	gs.CreatedAt = fromDb.CreatedAt.Time
	gs.UpdatedAt = &fromDb.UpdatedAt.Time
	return nil
}

func (gs *GameStatistic) ModelToAPI() oapipublic.GameStatistic {
	formatted := gs.FormatStatistics()

	return oapipublic.GameStatistic{
		AveragePerPoint:        formatted.AverageTimePerPointSeconds,
		LeftConsecutivePoints:  formatted.LeftConsecutivePoints,
		RightConsecutivePoints: formatted.RightConsecutivePoints,
		LongestPoint:           formatted.LongestPointSeconds,
		LongestPointTeam:       formatted.LongestPointTeam,
		ShortestPoint:          formatted.ShortestPointSeconds,
		ShortestPointTeam:      formatted.ShortestPointTeam,
		TotalGameTime:          formatted.TotalGameTimeSeconds,
		LeftAveragePerPoint:    formatted.LeftAverageTimePerPointSeconds,
		RightAveragePerPoint:   formatted.RightAverageTimePerPointSeconds,
	}
}

func (gs GameStatistic) FormatStatistics() FormattedGameStatistic {
	return FormattedGameStatistic{
		AverageTimePerPointSeconds:      fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.AverageTimePerPointSeconds)/60), gs.AverageTimePerPointSeconds%60),
		LeftConsecutivePoints:           strconv.Itoa(gs.LeftConsecutivePoints),
		RightConsecutivePoints:          strconv.Itoa(gs.RightConsecutivePoints),
		LongestPointSeconds:             fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.LongestPointSeconds)/60), gs.LongestPointSeconds%60),
		LongestPointTeam:                gs.LongestPointTeam,
		ShortestPointSeconds:            fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.ShortestPointSeconds)/60), gs.ShortestPointSeconds%60),
		ShortestPointTeam:               gs.ShortestPointTeam,
		TotalGameTimeSeconds:            fmt.Sprintf("%0.f h %d m", math.Floor(float64(gs.TotalGameTimeSeconds)/60/60), (gs.TotalGameTimeSeconds/60)%60),
		RightAverageTimePerPointSeconds: fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.RightAverageTimePerPointSeconds)/60), gs.RightAverageTimePerPointSeconds%60),
		LeftAverageTimePerPointSeconds:  fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.LeftAverageTimePerPointSeconds)/60), gs.LeftAverageTimePerPointSeconds%60),
	}
}
