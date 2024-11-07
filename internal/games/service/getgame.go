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

	if strings.Contains(previousServer, serverCompareString) {
		streakPoints = 0
	}

	streakPoints += 1
	updatedStreakPoints = streakPoints

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

	var longestPoint, shortestPoint, totalSeconds int
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

			if timeDiff > longestPoint {
				longestPoint = timeDiff
			}
			if timeDiff < shortestPoint || shortestPoint == 0 {
				shortestPoint = timeDiff
			}
		}
	}

	averageSeconds := totalSeconds / len(apiSteps)

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
			TotalGameTime:          game.GetGameLength(gameSteps[len(gameSteps)-1].ScoreAt),
			RightConsecutivePoints: strconv.Itoa(rightConsecutivePoints),
			LeftConsecutivePoints:  strconv.Itoa(leftConsecutivePoints),
			LongestPoint:           fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(longestPoint)/60), longestPoint%60),
			ShortestPoint:          fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(shortestPoint)/60), shortestPoint%60),
			AveragePerPoint:        fmt.Sprintf("%02.fm %02.ds", math.Floor(float64(averageSeconds)/60), averageSeconds%60),
		},
	}, nil
}
