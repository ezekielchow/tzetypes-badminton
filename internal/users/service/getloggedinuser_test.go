package users

import (
	"common/models"
	"common/oapiprivate"
	"common/utils"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLoggedInUser(t *testing.T) {
	ctx := context.Background()
	userService := InitService(ctx)

	t.Run("successful with no updated at", func(t *testing.T) {
		user := models.User{
			ID:           "UNIQUE_ID",
			Email:        utils.NewEmail(10),
			PasswordHash: "",
			CreatedAt:    time.Now(),
			UpdatedAt:    nil,
		}

		res, err := userService.GetLoggedInUser(ctx, oapiprivate.GetLoggedInUserRequestObject{}, user)
		assert.NoError(t, err)

		successRes, ok := res.(oapiprivate.GetLoggedInUser200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, oapiprivate.User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: "",
		}, successRes.User)
	})

	t.Run("successful with updated at", func(t *testing.T) {

		now := time.Now()
		user := models.User{
			ID:           "UNIQUE_ID",
			Email:        utils.NewEmail(10),
			PasswordHash: "",
			CreatedAt:    now,
			UpdatedAt:    &now,
		}

		res, err := userService.GetLoggedInUser(ctx, oapiprivate.GetLoggedInUserRequestObject{}, user)
		assert.NoError(t, err)

		successRes, ok := res.(oapiprivate.GetLoggedInUser200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, oapiprivate.User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		}, successRes.User)
	})

}
