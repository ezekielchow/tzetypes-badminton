package users

import (
	"common/oapipublic"
	"common/utils"
	"context"
	"testing"

	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {

	ctx := context.Background()

	userService := InitService(ctx)

	t.Run("Successful signup", func(t *testing.T) {
		email := utils.NewEmail(10)

		res, err := userService.Signup(ctx, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				Email:          types.Email(email),
				Password:       "",
				PasswordRepeat: "",
			},
		})
		assert.NoError(t, err)
		_, ok := res.(oapipublic.Signup201Response)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		foundUser, err := userService.UserStore.FindUserWithEmail(ctx, nil, email)
		if err != nil {
			t.Fatalf("unable to find user: %s", err)
		}

		if foundUser.ID == "" {
			t.Fatal("failed to create user")
		}

		foundClub, err := userService.ClubStore.GetClubGivenOwnerId(ctx, nil, foundUser.ID)
		if err != nil {
			t.Fatalf("unable to find club: %s", err)
		}
		if foundClub.ID == "" {
			t.Fatal("failed to create club")
		}

	})

	t.Run("repeat password does not match password", func(t *testing.T) {
		email := utils.NewEmail(10)

		res, err := userService.Signup(ctx, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				Email:          types.Email(email),
				Password:       "123",
				PasswordRepeat: "1234",
			},
		})

		assert.NoError(t, err)

		defaultResponse, ok := res.(oapipublic.SignupdefaultJSONResponse)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		assert.Equal(t, PasswordRepeatError, defaultResponse.Body.Message)
	})

	t.Run("email used error", func(t *testing.T) {

		email := utils.NewEmail(10)

		created, err := userService.UserStore.CreateUser(ctx, nil, email, "")
		if err != nil {
			t.Fatalf("unable to create user: %s", err)
		}

		res, err := userService.Signup(ctx, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				Email:          types.Email(created.Email),
				Password:       "123",
				PasswordRepeat: "123",
			},
		})

		assert.NoError(t, err)

		defaultResponse, ok := res.(oapipublic.SignupdefaultJSONResponse)
		if !ok {
			t.Fatal("unable to convert to default response")
		}

		assert.Equal(t, EmailUsedError, defaultResponse.Body.Message)
	})
}
