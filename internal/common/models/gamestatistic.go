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
	LeftLongestPointSeconds         int
	LeftShortestPointSeconds        int
	RightLongestPointSeconds        int
	RightShortestPointSeconds       int
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
	LeftLongestPoint                string
	LeftShortestPoint               string
	RightLongestPoint               string
	RightShortestPoint              string
	AverageTimePerPointSeconds      string
	LeftAverageTimePerPointSeconds  string
	RightAverageTimePerPointSeconds string
	ConsecutivePointsRatio          string
	LongestPointRatio               string
	ShortestPointRatio              string
	AverageTimePerPointRatio        string
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
	gs.LeftLongestPointSeconds = int(*fromDb.LeftLongestPointSeconds)
	gs.LeftShortestPointSeconds = int(*fromDb.LeftShortestPointSeconds)
	gs.RightLongestPointSeconds = int(*fromDb.RightLongestPointSeconds)
	gs.RightShortestPointSeconds = int(*fromDb.RightShortestPointSeconds)
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
		LeftLongestPoint:       formatted.LeftLongestPoint,
		LeftShortestPoint:      formatted.LeftShortestPoint,
		RightLongestPoint:      formatted.RightLongestPoint,
		RightShortestPoint:     formatted.RightShortestPoint,
		TotalGameTime:          formatted.TotalGameTimeSeconds,
		LeftAveragePerPoint:    formatted.LeftAverageTimePerPointSeconds,
		RightAveragePerPoint:   formatted.RightAverageTimePerPointSeconds,
		ConsecutivePointsRatio: formatted.ConsecutivePointsRatio,
		LongestPointRatio:      formatted.LongestPointRatio,
		ShortestPointRatio:     formatted.ShortestPointRatio,
		AveragePerPointRatio:   formatted.AverageTimePerPointRatio,
	}
}

func calculateRatios(left, right int) (leftRatio, rightRatio float64) {
	total := left + right
	if total != 0 {
		leftRatio = (float64(left) / float64(total)) * 100
		rightRatio = 100 - leftRatio
	} else {
		leftRatio = 100
		rightRatio = 0
	}
	return
}

func (gs GameStatistic) FormatStatistics() FormattedGameStatistic {

	leftConsecutiveRatio, rightConsecutiveRatio := calculateRatios(gs.LeftConsecutivePoints, gs.RightConsecutivePoints)
	leftLongestRatio, rightLongestRatio := calculateRatios(gs.LeftLongestPointSeconds, gs.RightLongestPointSeconds)
	leftShortestRatio, rightShortestRatio := calculateRatios(gs.LeftShortestPointSeconds, gs.RightShortestPointSeconds)
	leftAveragePerPointRatio, rightAveragePerPointRatio := calculateRatios(gs.LeftAverageTimePerPointSeconds, gs.RightAverageTimePerPointSeconds)

	return FormattedGameStatistic{
		AverageTimePerPointSeconds:      fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.AverageTimePerPointSeconds)/60), gs.AverageTimePerPointSeconds%60),
		LeftConsecutivePoints:           strconv.Itoa(gs.LeftConsecutivePoints),
		RightConsecutivePoints:          strconv.Itoa(gs.RightConsecutivePoints),
		LeftLongestPoint:                fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.LeftLongestPointSeconds)/60), gs.LeftLongestPointSeconds%60),
		LeftShortestPoint:               fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.LeftShortestPointSeconds)/60), gs.LeftShortestPointSeconds%60),
		RightLongestPoint:               fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.RightLongestPointSeconds)/60), gs.RightLongestPointSeconds%60),
		RightShortestPoint:              fmt.Sprintf("%0.fm %ds", math.Floor(float64(gs.RightShortestPointSeconds)/60), gs.RightShortestPointSeconds%60),
		TotalGameTimeSeconds:            fmt.Sprintf("%0.f h %d m", math.Floor(float64(gs.TotalGameTimeSeconds)/60/60), (gs.TotalGameTimeSeconds/60)%60),
		RightAverageTimePerPointSeconds: fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.RightAverageTimePerPointSeconds)/60), gs.RightAverageTimePerPointSeconds%60),
		LeftAverageTimePerPointSeconds:  fmt.Sprintf("%.fm %ds", math.Floor(float64(gs.LeftAverageTimePerPointSeconds)/60), gs.LeftAverageTimePerPointSeconds%60),
		ConsecutivePointsRatio:          fmt.Sprintf("%.2f", leftConsecutiveRatio) + ":" + fmt.Sprintf("%.2f", rightConsecutiveRatio),
		LongestPointRatio:               fmt.Sprintf("%.2f", leftLongestRatio) + ":" + fmt.Sprintf("%.2f", rightLongestRatio),
		ShortestPointRatio:              fmt.Sprintf("%.2f", leftShortestRatio) + ":" + fmt.Sprintf("%.2f", rightShortestRatio),
		AverageTimePerPointRatio:        fmt.Sprintf("%.2f", leftAveragePerPointRatio) + ":" + fmt.Sprintf("%.2f", rightAveragePerPointRatio),
	}
}
