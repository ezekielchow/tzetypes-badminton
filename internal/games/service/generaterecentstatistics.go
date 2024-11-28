package games

import (
	"common/models"
	"common/oapipublic"
	"context"
)

const (
	GenerateRecentStatisticsNotFoundError = "information not sufficient to generate statistics"
)

func getRecentStatistics(ctx context.Context, gs GameService, userId string) (models.GameRecentStatistic, error) {

	grs := models.GameRecentStatistic{}

	histories, err := gs.GameStore.GetMostRecentGameHistories(ctx, nil, userId)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	grs.NeedsRegenerating = 0
	grs.UserID = userId
	grs.GameCount = len(histories)

	totalTimePerPoints := 0
	totalTimePerPointsWon := 0
	totalTimePerPointsLost := 0
	grs.LongestRallySeconds = 0
	grs.ShortestRallySeconds = int(^uint(0) >> 1)

	for _, history := range histories {

		if history.IsGameWon == 1 {
			grs.Wins += 1
		} else {
			grs.Losses += 1
		}

		grs.TotalPoints += history.TotalPoints
		grs.PointsWon += history.PointsWon

		totalTimePerPoints += history.AverageTimePerPointSeconds
		totalTimePerPointsWon += history.AverageTimePerPointWonSeconds
		totalTimePerPointsLost += history.AverageTimePerPointLostSeconds

		if history.LongestRallySeconds > grs.LongestRallySeconds {
			grs.LongestRallySeconds = history.LongestRallySeconds
			grs.LongestRallyIsWon = history.LongestRallyIsWon
		}

		if history.ShortestRallySeconds < grs.ShortestRallySeconds {
			grs.ShortestRallySeconds = history.ShortestRallySeconds
			grs.ShortestRallyIsWon = history.ShortestRallyIsWon
		}
	}

	if totalTimePerPoints > 0 {
		grs.AverageTimePerPointSeconds = totalTimePerPoints / grs.GameCount
	}
	if totalTimePerPointsWon > 0 {
		grs.AverageTimePerPointWonSeconds = totalTimePerPointsWon / grs.GameCount
	}
	if totalTimePerPointsLost > 0 {
		grs.AverageTimePerPointLostSeconds = totalTimePerPointsLost / grs.GameCount
	}

	return grs, nil
}

func (gs GameService) GenerateRecentStatistics(ctx context.Context, input oapipublic.GenerateRecentStatisticsRequestObject) (oapipublic.GenerateRecentStatisticsResponseObject, error) {

	// limits to 10 users
	needsGeneration, err := gs.GameStore.GetGameRecentStatisticThatNeedsRegeneration(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, toGenerate := range needsGeneration {
		recentStatistic, err := getRecentStatistics(ctx, gs, toGenerate.UserID)
		if err != nil {
			return nil, err
		}

		_, err = gs.GameStore.CreateOrUpdateGameRecentStatistic(ctx, nil, recentStatistic)
		if err != nil {
			return nil, err
		}
	}

	return oapipublic.GenerateRecentStatistics200Response{}, nil
}
