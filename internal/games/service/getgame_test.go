package games

import (
	"common/models"
	"common/oapipublic"
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGame(t *testing.T) {

	gameService := InitService(context.Background())

	t.Run("statistic returning correct values", func(t *testing.T) {
		ctx := context.Background()

		clubToCreate := models.ClubFactory(1, map[string]interface{}{})[0]
		clubToCreate.ID = ""
		club, err := gameService.ClubStore.CreateClub(ctx, nil, clubToCreate)
		if err != nil {
			t.Errorf("unable to create club: %s", err.Error())
		}

		gameToCreate := models.GameFactory(1, map[string]interface{}{
			"ClubID": club.ID,
		})[0]
		game, err := gameService.GameStore.CreateGame(ctx, nil, gameToCreate)
		if err != nil {
			t.Errorf("unable to create game: %s", err.Error())
		}

		initialTime := time.Now()
		secondLongestPoint := initialTime.Add(time.Minute * 2)
		thirdShortestPoint := secondLongestPoint.Add(time.Second * 2)
		fourthPointTime := thirdShortestPoint.Add(time.Second * 15)

		toCreateSteps := []map[string]interface{}{
			{
				"GameID":         game.ID,
				"TeamLeftScore":  0,
				"TeamRightScore": 0,
				"ScoreAt":        initialTime,
				"StepNum":        1,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  1,
				"TeamRightScore": 0,
				"ScoreAt":        secondLongestPoint,
				"StepNum":        2,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 0,
				"ScoreAt":        thirdShortestPoint,
				"StepNum":        3,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 1,
				"ScoreAt":        fourthPointTime,
				"StepNum":        4,
			},
		}

		steps := []models.GameStep{}
		for _, step := range toCreateSteps {
			created, err := gameService.GameStore.CreateGameStep(ctx, nil, models.GameStepFactory(1, step)[0])
			if err != nil {
				t.Errorf("unable to create game step: %s", err.Error())
			}

			steps = append(steps, created)
		}

		statistic, err := generateGameStatistics(steps)
		if err != nil {
			t.Errorf("unable to create statistic: %s", err.Error())
		}

		assert.Equal(t, int(fourthPointTime.Sub(initialTime).Seconds()), statistic.TotalGameTimeSeconds)
		assert.Equal(t, 1, statistic.RightConsecutivePoints)
		assert.Equal(t, 2, statistic.LeftConsecutivePoints)
		assert.Equal(t, int(secondLongestPoint.Sub(initialTime).Seconds()), statistic.LeftLongestPointSeconds)
		assert.Equal(t, int(thirdShortestPoint.Sub(secondLongestPoint).Seconds()), statistic.LeftShortestPointSeconds)
		assert.Equal(t, int(fourthPointTime.Sub(thirdShortestPoint).Seconds()), statistic.RightLongestPointSeconds)
		assert.Equal(t, int(fourthPointTime.Sub(thirdShortestPoint).Seconds()), statistic.RightShortestPointSeconds)
		assert.Equal(t, int((fourthPointTime.Sub(initialTime).Seconds())/4), statistic.AverageTimePerPointSeconds)
		assert.Equal(t, 122/2, statistic.LeftAverageTimePerPointSeconds)
		assert.Equal(t, 15, statistic.RightAverageTimePerPointSeconds)
	})

	t.Run("get game successful", func(t *testing.T) {
		ctx := context.Background()

		clubToCreate := models.ClubFactory(1, map[string]interface{}{})[0]
		clubToCreate.ID = ""
		club, err := gameService.ClubStore.CreateClub(ctx, nil, clubToCreate)
		if err != nil {
			t.Errorf("unable to create club: %s", err.Error())
		}

		gameToCreate := models.GameFactory(1, map[string]interface{}{
			"ClubID": club.ID,
		})[0]
		game, err := gameService.GameStore.CreateGame(ctx, nil, gameToCreate)
		if err != nil {
			t.Errorf("unable to create game: %s", err.Error())
		}

		initialTime := time.Now()
		time2 := initialTime.Add(time.Minute * 2)
		time3 := time2.Add(time.Second * 2)
		time4 := time3.Add(time.Second * 15)
		time5 := time4.Add(time.Second * 10)
		time6 := time5.Add(time.Minute * 1)

		toCreateSteps := []map[string]interface{}{
			{
				"GameID":         game.ID,
				"TeamLeftScore":  0,
				"TeamRightScore": 0,
				"ScoreAt":        initialTime,
				"StepNum":        1,
				"CurrentServer":  models.LeftEvenServer,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  1,
				"TeamRightScore": 0,
				"ScoreAt":        time2,
				"StepNum":        2,
				"CurrentServer":  models.LeftOddServer,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 0,
				"ScoreAt":        time3,
				"StepNum":        3,
				"CurrentServer":  models.LeftEvenServer,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 1,
				"ScoreAt":        time4,
				"StepNum":        4,
				"CurrentServer":  models.RightOddServer,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 2,
				"ScoreAt":        time5,
				"StepNum":        5,
				"CurrentServer":  models.RightEvenServer,
			},
			{
				"GameID":         game.ID,
				"TeamLeftScore":  2,
				"TeamRightScore": 3,
				"ScoreAt":        time6,
				"StepNum":        6,
				"CurrentServer":  models.RightOddServer,
			},
		}

		for _, step := range toCreateSteps {
			_, err := gameService.GameStore.CreateGameStep(ctx, nil, models.GameStepFactory(1, step)[0])
			if err != nil {
				t.Errorf("unable to create game step: %s", err.Error())
			}
		}

		res, err := gameService.GetGame(ctx, oapipublic.GetGameRequestObject{
			GameId: game.ID,
		})
		if err != nil {
			t.Fatalf("unagle to get game: %s", err.Error())
		}

		resSuccess, ok := res.(oapipublic.GetGame200JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, game.ID, resSuccess.Game.Id)
		assert.Equal(t, 6, len(resSuccess.Steps))

		perPoint := time6.Sub(initialTime).Seconds() / 6
		assert.Equal(t, fmt.Sprintf("%.fm %ds", math.Floor(perPoint/60), int(perPoint)%60), resSuccess.Statistics.AveragePerPoint)
		assert.Equal(t, "2", resSuccess.Statistics.LeftConsecutivePoints)
		assert.Equal(t, "3", resSuccess.Statistics.RightConsecutivePoints)

		leftLongestPointSeconds := time2.Sub(initialTime).Seconds()
		assert.Equal(t, fmt.Sprintf("%0.fm %ds", math.Floor(leftLongestPointSeconds/60), int(leftLongestPointSeconds)%60), resSuccess.Statistics.LeftLongestPoint)

		leftShortestPointSeconds := time3.Sub(time2).Seconds()
		assert.Equal(t, fmt.Sprintf("%0.fm %ds", math.Floor(leftShortestPointSeconds)/60, int(leftShortestPointSeconds)%60), resSuccess.Statistics.LeftShortestPoint)

		rightLongestPointSeconds := time6.Sub(time5).Seconds()
		assert.Equal(t, fmt.Sprintf("%0.fm %ds", math.Floor(rightLongestPointSeconds/60), int(rightLongestPointSeconds)%60), resSuccess.Statistics.RightLongestPoint)

		rightShortestPointSeconds := time5.Sub(time4).Seconds()
		assert.Equal(t, fmt.Sprintf("%0.fm %ds", math.Floor(rightShortestPointSeconds/60), int(rightShortestPointSeconds)%60), resSuccess.Statistics.RightShortestPoint)

		totalGameTime := time6.Sub(initialTime).Seconds()
		assert.Equal(t, fmt.Sprintf("%0.f h %d m", math.Floor(totalGameTime/60/60), (int(totalGameTime)/60)%60), resSuccess.Statistics.TotalGameTime)

		assert.Equal(t, "0m 28s", resSuccess.Statistics.RightAveragePerPoint)

		leftAverageTime := time3.Sub(initialTime).Seconds() / 2
		assert.Equal(t, fmt.Sprintf("%.fm %ds", math.Floor(leftAverageTime/60), int(leftAverageTime)%60), resSuccess.Statistics.LeftAveragePerPoint)

		consecutivePointsRatio := "40.00:60.00"
		assert.Equal(t, consecutivePointsRatio, resSuccess.Statistics.ConsecutivePointsRatio)

		longestPointRatio := "66.67:33.33"
		assert.Equal(t, longestPointRatio, resSuccess.Statistics.LongestPointRatio)

		shortestPointRatio := "16.67:83.33"
		assert.Equal(t, shortestPointRatio, resSuccess.Statistics.ShortestPointRatio)

		averageTimePerPointRatio := "68.54:31.46"
		assert.Equal(t, averageTimePerPointRatio, resSuccess.Statistics.AveragePerPointRatio)
	})
}
