package userservice

import (
	"common/models"
	"common/oapipublic"
	"context"
	"mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_validateSignup(t *testing.T) {
	t.Run("No errors", func(t *testing.T) {
		t.Parallel()

		repoMock := mocks.NewUserRepository(t)
		repoMock.On("FindUserWithEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(func(ctx context.Context, email string) (models.User, error) {
				return models.User{}, nil
			})

		err := validateSignup(context.Background(), repoMock, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				Email:          "test@email.com",
				Password:       "abc123",
				PasswordRepeat: "abc123",
			},
		})

		assert.NoError(t, err)
	})

	t.Run("repeat password does not match password", func(t *testing.T) {
		t.Parallel()

		repoMock := mocks.NewUserRepository(t)

		err := validateSignup(context.Background(), repoMock, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				Email:          "test@email.com",
				Password:       "abc1234",
				PasswordRepeat: "abc123",
			},
		})

		assert.Equal(t, PasswordRepeatError, err.Error())
	})

	t.Run("repeat password does not match password", func(t *testing.T) {
		t.Parallel()

		repoMock := mocks.NewUserRepository(t)
		repoMock.On("FindUserWithEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(func(ctx context.Context, email string) (models.User, error) {
				user := models.User{}
				user.Mock()

				return user, nil
			})

		err := validateSignup(context.Background(), repoMock, oapipublic.SignupRequestObject{
			Body: &oapipublic.SignupJSONRequestBody{
				// Email wont match because seed email is only letters
				Email:          "123@email.com",
				Password:       "abc123",
				PasswordRepeat: "abc123",
			},
		})

		assert.Equal(t, EmailUsedError, err.Error())
	})
}
