package users

import (
	"common/oapipublic"
	"common/utils"
	"context"
	"testing"

	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	userService := InitService(ctx)

	t.Run("sucessful login", func(t *testing.T) {
		email := utils.NewEmail(10)
		password := utils.NewString(15)

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			t.Fatalf("unable to generate password: %s", err)
		}

		_, err = userService.UserStore.CreateUser(ctx, nil, email, string(hash))
		if err != nil {
			t.Fatalf("unable to create user: %s", err)
		}

		res, err := userService.Login(ctx, oapipublic.LoginRequestObject{
			Body: &oapipublic.LoginJSONRequestBody{
				Email:    types.Email(email),
				Password: password,
			},
		})
		assert.NoError(t, err)

		successRes, ok := res.(Login200JSONResponse)
		if !ok {
			t.Fatalf("unable to convert response: %s", err)
		}

		assert.NotNil(t, successRes.RefreshToken)
		assert.NotNil(t, successRes.SessionToken)

		session, err := userService.SessionStore.FindSessionWithSessionID(ctx, nil, successRes.SessionToken)
		assert.NoError(t, err)

		assert.Equal(t, session.SessionToken, successRes.SessionToken)
	})

	t.Run("wrong email", func(t *testing.T) {
		email := utils.NewEmail(10)
		password := utils.NewString(15)

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			t.Fatalf("unable to generate password: %s", err)
		}

		_, err = userService.UserStore.CreateUser(ctx, nil, email, string(hash))
		if err != nil {
			t.Fatalf("unable to create user: %s", err)
		}

		wrongEmail := utils.NewEmail(10)

		res, err := userService.Login(ctx, oapipublic.LoginRequestObject{
			Body: &oapipublic.LoginJSONRequestBody{
				Email:    types.Email(wrongEmail),
				Password: password,
			},
		})
		assert.NoError(t, err)

		successRes, ok := res.(oapipublic.LogindefaultJSONResponse)
		if !ok {
			t.Fatalf("unable to convert response: %s", err)
		}

		assert.Equal(t, LoginCredentialsInvalidError, successRes.Body.Message)
	})

	t.Run("wrong email", func(t *testing.T) {
		email := utils.NewEmail(10)
		password := utils.NewString(15)

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			t.Fatalf("unable to generate password: %s", err)
		}

		_, err = userService.UserStore.CreateUser(ctx, nil, email, string(hash))
		if err != nil {
			t.Fatalf("unable to create user: %s", err)
		}

		res, err := userService.Login(ctx, oapipublic.LoginRequestObject{
			Body: &oapipublic.LoginJSONRequestBody{
				Email:    types.Email(email),
				Password: "WRONG PASSWORD",
			},
		})
		assert.NoError(t, err)

		successRes, ok := res.(oapipublic.LogindefaultJSONResponse)
		if !ok {
			t.Fatalf("unable to convert response: %s", err)
		}

		assert.Equal(t, LoginCredentialsInvalidError, successRes.Body.Message)
	})
}
