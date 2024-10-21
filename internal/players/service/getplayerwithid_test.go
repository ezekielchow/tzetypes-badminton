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

func TestGetPlayerWithId(t *testing.T) {

	ctx := context.Background()
	playerService := InitService(ctx)

	t.Run("get player successfully", func(t *testing.T) {

		ct := context.Background()
		name := utils.NewString(10)

		created, err := playerService.PlayerStore.CreatePlayer(ct, nil, models.Player{
			UserID: uuid.NewString(),
			Name:   name,
		}, "")
		if err != nil {
			t.Fatalf("unable to create player: %s", err)
		}

		res, err := playerService.GetPlayerWithId(ct, oapiprivate.GetPlayerWithIdRequestObject{
			Id: created.ID,
		})
		if err != nil {
			t.Fatalf("unable to get player: %s", err)
		}

		successRes, ok := res.(oapiprivate.GetPlayerWithId200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, created.Name, successRes.Name)
	})

	t.Run("get player with doesnt exist id", func(t *testing.T) {
		ct := context.Background()

		res, err := playerService.GetPlayerWithId(ct, oapiprivate.GetPlayerWithIdRequestObject{
			Id: uuid.NewString(),
		})
		assert.Nil(t, res)
		assert.Equal(t, err.Error(), GetPlayerWithIdNotFoundError)
	})
}
