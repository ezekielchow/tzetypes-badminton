package players

import (
	"common/models"
	"common/oapiprivate"
	"common/utils"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePlayer(t *testing.T) {
	playerService := InitService(context.Background())

	t.Run("update player successfully", func(t *testing.T) {
		ctx := context.Background()

		name := utils.NewString(10)

		created, err := playerService.PlayerStore.CreatePlayer(ctx, nil, models.Player{
			UserID: uuid.NewString(),
			Name:   name,
		}, "")
		if err != nil {
			t.Fatalf("unable to create player: %s", err)
		}

		toUpdateName := utils.NewString(10)
		res, err := playerService.UpdatePlayer(ctx, oapiprivate.UpdatePlayerWithIdRequestObject{
			Id: created.ID,
			Body: &oapiprivate.UpdatePlayerWithIdJSONRequestBody{
				Name: toUpdateName,
			},
		})
		if err != nil {
			t.Fatalf("unable to update player: %s", err)
		}

		successRes, ok := res.(oapiprivate.UpdatePlayerWithId200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		resGet, err := playerService.GetPlayerWithId(ctx, oapiprivate.GetPlayerWithIdRequestObject{
			Id: created.ID,
		})
		if err != nil {
			t.Fatalf("unable to get player: %s", err)
		}

		updatedPlayerRes, ok := resGet.(oapiprivate.GetPlayerWithId200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, successRes.Name, updatedPlayerRes.Name)
	})
}
