package users

import (
	"common/models"
	"common/oapipublic"
	"common/utils"
	"context"
	"testing"

	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestSignupPlayer(t *testing.T) {

	ctx := context.Background()

	userService := InitService(ctx)

	t.Run("Successful player signup", func(t *testing.T) {
		email := utils.NewEmail(10)

		res, err := userService.SignupPlayer(ctx, oapipublic.SignupPlayerRequestObject{
			Body: &oapipublic.SignupPlayerJSONRequestBody{
				Email:          types.Email(email),
				Password:       "",
				PasswordRepeat: "",
			},
		})
		assert.NoError(t, err)
		_, ok := res.(oapipublic.SignupPlayer201Response)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		foundUser, err := userService.UserStore.FindUserWithEmail(ctx, nil, email)
		if err != nil {
			t.Fatalf("unable to find user: %s", err)
		}
		assert.Equal(t, string(models.UserTypePlayer), foundUser.UserType, "user check for player")

		if foundUser.ID == "" {
			t.Fatal("failed to create user")
		}
	})

	t.Run("repeat password does not match password", func(t *testing.T) {
		email := utils.NewEmail(10)

		res, err := userService.SignupPlayer(ctx, oapipublic.SignupPlayerRequestObject{
			Body: &oapipublic.SignupPlayerJSONRequestBody{
				Email:          types.Email(email),
				Password:       "123",
				PasswordRepeat: "1234",
			},
		})

		assert.NoError(t, err)

		defaultResponse, ok := res.(oapipublic.SignupPlayerdefaultJSONResponse)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		assert.Equal(t, PasswordRepeatError, defaultResponse.Body.Message)
	})

	t.Run("email used error", func(t *testing.T) {

		email := utils.NewEmail(10)

		created, err := userService.UserStore.CreateUser(ctx, nil, email, "", string(models.UserTypePlayer))
		if err != nil {
			t.Fatalf("unable to create user: %s", err)
		}

		res, err := userService.SignupPlayer(ctx, oapipublic.SignupPlayerRequestObject{
			Body: &oapipublic.SignupPlayerJSONRequestBody{
				Email:          types.Email(created.Email),
				Password:       "123",
				PasswordRepeat: "123",
			},
		})

		assert.NoError(t, err)

		defaultResponse, ok := res.(oapipublic.SignupPlayerdefaultJSONResponse)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		assert.Equal(t, EmailUsedError, defaultResponse.Body.Message)
	})
}
