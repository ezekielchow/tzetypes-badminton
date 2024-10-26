package games

import (
	"common/models"
	"common/oapiprivate"
	"common/utils"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStartGame(t *testing.T) {

	gameService := InitService(context.Background())

	t.Run("successfully start game singles", func(t *testing.T) {

		ctx := context.Background()

		userID := uuid.New()

		club, err := gameService.ClubStore.CreateClub(ctx, nil, models.Club{
			OwnerID: userID.String(),
			Name:    utils.NewString(10),
		})
		if err != nil {
			t.Fatalf("unable to create club:%s", err)
		}

		leftEvenPlayerName := utils.NewString(10)
		rightEvenPlayerName := utils.NewString(10)

		res, err := gameService.StartGame(ctx, oapiprivate.StartGameRequestObject{
			Body: &oapiprivate.StartGameJSONRequestBody{
				GameType:            oapiprivate.Singles,
				LeftEvenPlayerName:  leftEvenPlayerName,
				RightEvenPlayerName: rightEvenPlayerName,
				ServingSide:         oapiprivate.RightEven,
			},
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("failed to start game:%s", err)
		}

		resSuccess, ok := res.(oapiprivate.StartGame201JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response:%s", err)
		}

		assert.Equal(t, club.ID, resSuccess.Game.ClubId)
		assert.Equal(t, string(oapiprivate.Singles), resSuccess.Game.GameType)
		assert.Equal(t, string(oapiprivate.RightEven), resSuccess.Game.ServingSide)
		assert.Equal(t, leftEvenPlayerName, resSuccess.Game.LeftEvenPlayerName)
		assert.Equal(t, rightEvenPlayerName, resSuccess.Game.RightEvenPlayerName)
		assert.Equal(t, "", resSuccess.Game.RightOddPlayerName)
		assert.Equal(t, "", resSuccess.Game.LeftOddPlayerName)
		assert.Equal(t, false, resSuccess.Game.IsEnded)
	})

	t.Run("successfully start game dobules", func(t *testing.T) {

		ctx := context.Background()

		userID := uuid.New()

		club, err := gameService.ClubStore.CreateClub(ctx, nil, models.Club{
			OwnerID: userID.String(),
			Name:    utils.NewString(10),
		})
		if err != nil {
			t.Fatalf("unable to create club:%s", err)
		}

		leftEvenPlayerName := utils.NewString(10)
		leftOddPlayerName := utils.NewString(10)
		rightEvenPlayerName := utils.NewString(10)
		rightOddPlayerName := utils.NewString(10)

		res, err := gameService.StartGame(ctx, oapiprivate.StartGameRequestObject{
			Body: &oapiprivate.StartGameJSONRequestBody{
				GameType:            oapiprivate.Doubles,
				LeftEvenPlayerName:  leftEvenPlayerName,
				LeftOddPlayerName:   &leftOddPlayerName,
				RightEvenPlayerName: rightEvenPlayerName,
				RightOddPlayerName:  &rightOddPlayerName,
				ServingSide:         oapiprivate.LeftEven,
			},
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("failed to start game:%s", err)
		}

		resSuccess, ok := res.(oapiprivate.StartGame201JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response:%s", err)
		}

		assert.Equal(t, club.ID, resSuccess.Game.ClubId)
		assert.Equal(t, string(oapiprivate.Doubles), resSuccess.Game.GameType)
		assert.Equal(t, string(oapiprivate.LeftEven), resSuccess.Game.ServingSide)
		assert.Equal(t, leftEvenPlayerName, resSuccess.Game.LeftEvenPlayerName)
		assert.Equal(t, leftOddPlayerName, resSuccess.Game.LeftOddPlayerName)
		assert.Equal(t, rightEvenPlayerName, resSuccess.Game.RightEvenPlayerName)
		assert.Equal(t, rightOddPlayerName, resSuccess.Game.RightOddPlayerName)
		assert.Equal(t, false, resSuccess.Game.IsEnded)
	})
}
