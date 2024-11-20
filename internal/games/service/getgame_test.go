package games

import (
	"common/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGame(t *testing.T) {

	gameService := InitService(context.Background())

	t.Run("Getting game statistics twice", func(t *testing.T) {
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
}
