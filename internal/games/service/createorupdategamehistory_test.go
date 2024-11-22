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

func CreateOrUpdateGameHistoryTest(t *testing.T) {
	gameService := InitService(context.Background())

	t.Run("successful create game history", func(t *testing.T) {
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

		res, err := gameService.CreateOrUpdateGameHistory(ctx, oapiprivate.CreateOrUpdateGameHistoryRequestObject{
			GameId: game.ID,
			Body: &oapiprivate.CreateOrUpdateGameHistoryRequestSchema{
				PlayerPosition: oapiprivate.LeftEvenPlayer,
			},
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("unable to create game:%s", err)
		}

		resSuccess, ok := res.(oapiprivate.CreateOrUpdateGameHistory200JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, userID.String(), resSuccess.GameHistory.UserId)
		assert.Equal(t, game.ID, resSuccess.GameHistory.GameId)
		assert.Equal(t, string(oapiprivate.LeftEvenPlayer), resSuccess.GameHistory.PlayerPosition)
	})

	t.Run("successful create, then update game history", func(t *testing.T) {
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

		res, err := gameService.CreateOrUpdateGameHistory(ctx, oapiprivate.CreateOrUpdateGameHistoryRequestObject{
			GameId: game.ID,
			Body: &oapiprivate.CreateOrUpdateGameHistoryRequestSchema{
				PlayerPosition: oapiprivate.LeftEvenPlayer,
			},
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("unable to create game:%s", err)
		}

		resSuccess, ok := res.(oapiprivate.CreateOrUpdateGameHistory200JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, userID.String(), resSuccess.GameHistory.UserId)
		assert.Equal(t, game.ID, resSuccess.GameHistory.GameId)
		assert.Equal(t, string(oapiprivate.LeftEvenPlayer), resSuccess.GameHistory.PlayerPosition)

		res, err = gameService.CreateOrUpdateGameHistory(ctx, oapiprivate.CreateOrUpdateGameHistoryRequestObject{
			GameId: game.ID,
			Body: &oapiprivate.CreateOrUpdateGameHistoryRequestSchema{
				PlayerPosition: oapiprivate.RightEvenPlayer,
			},
		}, models.User{
			ID: userID.String(),
		})
		if err != nil {
			t.Fatalf("unable to create game:%s", err)
		}

		resSuccess, ok = res.(oapiprivate.CreateOrUpdateGameHistory200JSONResponse)
		if !ok {
			t.Fatalf("failed to convert response")
		}

		assert.Equal(t, userID.String(), resSuccess.GameHistory.UserId)
		assert.Equal(t, game.ID, resSuccess.GameHistory.GameId)
		assert.Equal(t, string(oapiprivate.RightEvenPlayer), resSuccess.GameHistory.PlayerPosition)
		assert.NotNil(t, resSuccess.GameHistory.UpdatedAt)
	})
}
