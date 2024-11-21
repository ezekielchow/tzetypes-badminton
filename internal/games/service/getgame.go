package games

import (
	"common/models"
	"common/oapipublic"
	"context"
	"database/sql"
	"strings"
)

func updateStreak(currentTeam models.TeamSide, previousServer string, streakPoints int, leftConsecutivePoints int, rightConsecutivePoints int) (updatedStreakPoints int, updatedLeftConsecutivePoints int, updatedRightConsecutivePoints int) {

	serverCompareString := "right"
	if currentTeam == models.TeamSideRight {
		serverCompareString = "left"
	}

	// Reset streak points if the server side has changed
	if strings.Contains(previousServer, serverCompareString) {
		streakPoints = 0
	}

	// Increment the streak points for the current team
	streakPoints += 1
	updatedStreakPoints = streakPoints

	// Update consecutive points if the current streak is the highest for that team
	updatedLeftConsecutivePoints = leftConsecutivePoints
	updatedRightConsecutivePoints = rightConsecutivePoints

	if currentTeam == models.TeamSideLeft && updatedStreakPoints > leftConsecutivePoints {
		updatedLeftConsecutivePoints = updatedStreakPoints
	} else if currentTeam == models.TeamSideRight && updatedStreakPoints > rightConsecutivePoints {
		updatedRightConsecutivePoints = updatedStreakPoints
	}

	return updatedStreakPoints, updatedLeftConsecutivePoints, updatedRightConsecutivePoints
}

func generateGameStatistics(gameSteps []models.GameStep) (models.GameStatistic, error) {
	leftConsecutivePoints := 0
	rightConsecutivePoints := 0

	var longestLeftPointSeconds, shortestLeftPointSeconds int
	var longestRightPointSeconds, shortestRightPointSeconds int
	var totalSeconds, leftTotalSeconds, rightTotalSeconds int
	streakPoints := 0

	// Initialize shortest points with a high value
	const maxInt = int(^uint(0) >> 1)
	shortestLeftPointSeconds = maxInt
	shortestRightPointSeconds = maxInt

	apiSteps := []oapipublic.GameStep{}
	for i, step := range gameSteps {
		apiSteps = append(apiSteps, oapipublic.GameStep{
			CreatedAt:           step.CreatedAt.String(),
			GameId:              step.GameID,
			Id:                  step.ID,
			ScoreAt:             step.ScoreAt.String(),
			StepNum:             step.StepNum,
			TeamLeftScore:       step.TeamLeftScore,
			TeamRightScore:      step.TeamRightScore,
			CurrentServer:       step.CurrentServer,
			LeftEvenPlayerName:  step.LeftEvenPlayerName,
			LeftOddPlayerName:   *step.LeftOddPlayerName,
			RightEvenPlayerName: step.RightEvenPlayerName,
			RightOddPlayerName:  *step.RightOddPlayerName,
			UpdatedAt:           step.UpdatedAt.String(),
			SyncId:              &step.SyncId,
		})

		if i > 0 {
			previous := gameSteps[i-1]
			timeDiff := int(step.ScoreAt.Sub(previous.ScoreAt).Seconds())
			totalSeconds += timeDiff

			if step.TeamLeftScore > previous.TeamLeftScore {
				streakPoints, leftConsecutivePoints, rightConsecutivePoints = updateStreak(models.TeamSideLeft, previous.CurrentServer, streakPoints, leftConsecutivePoints, rightConsecutivePoints)
				leftTotalSeconds += timeDiff

				// Update longest and shortest point for Team Left
				if timeDiff > longestLeftPointSeconds {
					longestLeftPointSeconds = timeDiff
				}
				if timeDiff < shortestLeftPointSeconds {
					shortestLeftPointSeconds = timeDiff
				}
			} else {
				streakPoints, leftConsecutivePoints, rightConsecutivePoints = updateStreak(models.TeamSideRight, previous.CurrentServer, streakPoints, leftConsecutivePoints, rightConsecutivePoints)
				rightTotalSeconds += timeDiff

				// Update longest and shortest point for Team Right
				if timeDiff > longestRightPointSeconds {
					longestRightPointSeconds = timeDiff
				}
				if timeDiff < shortestRightPointSeconds {
					shortestRightPointSeconds = timeDiff
				}
			}
		}
	}

	// Avoid division by zero when calculating averages
	leftPoints := gameSteps[len(gameSteps)-1].TeamLeftScore
	rightPoints := gameSteps[len(gameSteps)-1].TeamRightScore

	leftAverageTimePerPoint := 0
	rightAverageTimePerPoint := 0
	if leftPoints > 0 {
		leftAverageTimePerPoint = leftTotalSeconds / leftPoints
	}
	if rightPoints > 0 {
		rightAverageTimePerPoint = rightTotalSeconds / rightPoints
	}

	totalGameTimeSeconds := int(gameSteps[len(gameSteps)-1].ScoreAt.Sub(gameSteps[0].ScoreAt).Seconds())

	return models.GameStatistic{
		TotalGameTimeSeconds:            totalGameTimeSeconds,
		RightConsecutivePoints:          rightConsecutivePoints,
		LeftConsecutivePoints:           leftConsecutivePoints,
		LeftLongestPointSeconds:         longestLeftPointSeconds,
		LeftShortestPointSeconds:        shortestLeftPointSeconds,
		RightLongestPointSeconds:        longestRightPointSeconds,
		RightShortestPointSeconds:       shortestRightPointSeconds,
		AverageTimePerPointSeconds:      totalSeconds / len(apiSteps),
		LeftAverageTimePerPointSeconds:  leftAverageTimePerPoint,
		RightAverageTimePerPointSeconds: rightAverageTimePerPoint,
	}, nil
}

func (gs GameService) GetGame(ctx context.Context, input oapipublic.GetGameRequestObject) (oapipublic.GetGameResponseObject, error) {

	game, err := gs.GameStore.GetGame(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	gameSteps, err := gs.GameStore.GetGameSteps(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	gameStatistics, err := gs.GameStore.GetStatisticsWithGameId(ctx, nil, input.GameId)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	if gameStatistics.ID != "" {
		apiSteps := models.GameStepsToAPI(gameSteps)
		formattedStatistics := gameStatistics.ModelToAPI()

		return oapipublic.GetGame200JSONResponse{
			Game:       game.ModelToPublicAPI(),
			Steps:      apiSteps,
			Statistics: &formattedStatistics,
		}, nil
	}

	gameStatistics, err = generateGameStatistics(gameSteps)
	if err != nil {
		return nil, err
	}

	createdStatistics, err := gs.GameStore.CreateStatistic(ctx, nil, game.ID, gameStatistics)
	if err != nil {
		return nil, err
	}

	formattedStatistics := createdStatistics.ModelToAPI()
	apiSteps := models.GameStepsToAPI(gameSteps)

	return oapipublic.GetGame200JSONResponse{
		Game:       game.ModelToPublicAPI(),
		Steps:      apiSteps,
		Statistics: &formattedStatistics,
	}, nil
}
