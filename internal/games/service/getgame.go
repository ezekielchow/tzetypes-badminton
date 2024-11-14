package games

import (
	"common/models"
	"common/oapipublic"
	"context"
	"fmt"
	"math"
	"strconv"
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

func (gs GameService) GetGame(ctx context.Context, input oapipublic.GetGameRequestObject) (oapipublic.GetGameResponseObject, error) {

	game, err := gs.GameStore.GetGame(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	gameSteps, err := gs.GameStore.GetGameSteps(ctx, nil, input.GameId)
	if err != nil {
		return nil, err
	}

	leftConsecutivePoints := 0
	rightConsecutivePoints := 0

	var longestPointSeconds, shortestPointSeconds, totalSeconds int
	var longestPointTeam, shortestPointTeam string
	streakPoints := 0
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
			// get streaks
			if step.TeamLeftScore > previous.TeamLeftScore {
				streakPoints, leftConsecutivePoints, rightConsecutivePoints = updateStreak(models.TeamSideLeft, previous.CurrentServer, streakPoints, leftConsecutivePoints, rightConsecutivePoints)
			} else {
				streakPoints, leftConsecutivePoints, rightConsecutivePoints = updateStreak(models.TeamSideRight, previous.CurrentServer, streakPoints, leftConsecutivePoints, rightConsecutivePoints)
			}

			// getscorediff
			timeDiff := int(step.ScoreAt.Sub(previous.ScoreAt).Seconds())
			totalSeconds += timeDiff

			if timeDiff > longestPointSeconds {
				longestPointSeconds = timeDiff
				if step.TeamLeftScore > previous.TeamLeftScore {
					longestPointTeam = string(models.TeamSideLeft)
				} else {
					longestPointTeam = string(models.TeamSideRight)
				}
			}
			if timeDiff < shortestPointSeconds || shortestPointSeconds == 0 {
				shortestPointSeconds = timeDiff
				if step.TeamLeftScore > previous.TeamLeftScore {
					shortestPointTeam = string(models.TeamSideLeft)
				} else {
					shortestPointTeam = string(models.TeamSideRight)
				}
			}
		}
	}

	averageSeconds := totalSeconds / len(apiSteps)
	totalGameTimeSeconds := game.GetGameLength(gameSteps[len(gameSteps)-1].ScoreAt)

	_, err = gs.GameStore.CreateStatistic(ctx, nil, game.ID, models.GameStatistic{
		TotalGameTimeSeconds:       totalGameTimeSeconds,
		RightConsecutivePoints:     rightConsecutivePoints,
		LeftConsecutivePoints:      leftConsecutivePoints,
		LongestPointSeconds:        longestPointSeconds,
		LongestPointTeam:           longestPointTeam,
		ShortestPointSeconds:       shortestPointSeconds,
		ShortestPointTeam:          shortestPointTeam,
		AverageTimePerPointSeconds: averageSeconds,
	})
	if err != nil {
		return nil, err
	}

	return oapipublic.GetGame200JSONResponse{
		Game: oapipublic.Game{
			ClubId:              game.ClubID,
			CreatedAt:           game.CreatedAt.String(),
			GameType:            game.GameType,
			Id:                  game.ID,
			LeftEvenPlayerName:  game.LeftEvenPlayerName,
			LeftOddPlayerName:   *game.LeftOddPlayerName,
			RightEvenPlayerName: game.RightEvenPlayerName,
			RightOddPlayerName:  *game.RightOddPlayerName,
			ServingSide:         game.ServingSide,
			IsEnded:             game.IsEnded,
			UpdatedAt:           game.UpdatedAt.String(),
		},
		Steps: apiSteps,
		Statistics: &oapipublic.GameStatistic{
			TotalGameTime:          models.GetGameLengthFormatted(totalGameTimeSeconds),
			RightConsecutivePoints: strconv.Itoa(rightConsecutivePoints),
			LeftConsecutivePoints:  strconv.Itoa(leftConsecutivePoints),
			LongestPoint:           fmt.Sprintf("%02.fm %ds", math.Floor(float64(longestPointSeconds)/60), longestPointSeconds%60),
			LongestPointTeam:       longestPointTeam,
			ShortestPoint:          fmt.Sprintf("%02.fm %ds", math.Floor(float64(shortestPointSeconds)/60), shortestPointSeconds%60),
			ShortestPointTeam:      shortestPointTeam,
			AveragePerPoint:        fmt.Sprintf("%02.fm %ds", math.Floor(float64(shortestPointSeconds)/60), averageSeconds%60),
		},
	}, nil
}
