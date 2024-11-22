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

func TestGetGameHistory(t *testing.T) {

	gameService := InitService(context.Background())

	t.Run("successfully get game player", func(t *testing.T) {

		ctx := context.Background()
		userID := uuid.New()

		club, err := gameService.ClubStore.CreateClub(ctx, nil, models.Club{
			OwnerID: userID.String(),
			Name:    utils.NewString(10),
		})
		if err != nil {
			t.Fatalf("unable to create club:%s", err)
		}

		game, err := gameService.GameStore.CreateGame(ctx, nil,
			models.GameFactory(1, map[string]interface{}{"ClubID": club.ID})[0])
		if err != nil {
			t.Fatalf("unable to create game:%s", err)
		}

		_, err = gameService.GameStore.CreateOrUpdateGameHistory(ctx, nil, models.GameHistory{
			UserID:         userID.String(),
			GameID:         game.ID,
			PlayerPosition: string(oapiprivate.LeftEvenPlayer),
		})
		if err != nil {
			t.Fatalf("unable to create game history:%s", err)
		}

		res, err := gameService.GetGameHistory(ctx, oapiprivate.GetGameHistoryRequestObject{
			GameId: game.ID,
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("unable to get game history:%s", err)
		}

		resSuccess, ok := res.(oapiprivate.GetGameHistory200JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, userID.String(), resSuccess.GameHistory.UserId)
		assert.Equal(t, game.ID, resSuccess.GameHistory.GameId)
		assert.Equal(t, string(oapiprivate.LeftEvenPlayer), resSuccess.GameHistory.PlayerPosition)
	})

	t.Run("returns not found error when history doesn't exist", func(t *testing.T) {
		ctx := context.Background()

		userID := uuid.New()

		club, err := gameService.ClubStore.CreateClub(ctx, nil, models.Club{
			OwnerID: userID.String(),
			Name:    utils.NewString(10),
		})
		if err != nil {
			t.Fatalf("unable to create club:%s", err)
		}

		game, err := gameService.GameStore.CreateGame(ctx, nil,
			models.GameFactory(1, map[string]interface{}{"ClubID": club.ID})[0])
		if err != nil {
			t.Fatalf("unable to create game:%s", err)
		}

		res, err := gameService.GetGameHistory(ctx, oapiprivate.GetGameHistoryRequestObject{
			GameId: game.ID,
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("unable to get game history:%s", err)
		}

		resNotFound, ok := res.(oapiprivate.GetGameHistorydefaultJSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, GetGameHistoryNotFoundError, resNotFound.Body.Message)
	})
}
