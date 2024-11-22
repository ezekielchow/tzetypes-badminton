package users

import (
	"common/models"
	"common/oapipublic"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"
	userstore "users/store"

	"golang.org/x/crypto/bcrypt"
)

const (
	SignupPlayerPasswordRepeatError = "repeated password has to be the same as password"
	SignupPlayerEmailUsedError      = "email is already registered with another account"
)

func returnSignupPlayerError(err error) oapipublic.SignupPlayerdefaultJSONResponse {
	return oapipublic.SignupPlayerdefaultJSONResponse{
		Body:       oapipublic.Error{Message: err.Error()},
		StatusCode: http.StatusInternalServerError,
	}
}

func validateSignupPlayer(ctx context.Context, userStore userstore.UserRepository, input oapipublic.SignupPlayerRequestObject) error {

	if input.Body.Password != input.Body.PasswordRepeat {
		return errors.New(PasswordRepeatError)
	}

	user, err := userStore.FindUserWithEmail(ctx, nil, string(input.Body.Email))
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return err
	}

	if user.ID != "" {
		return errors.New(EmailUsedError)
	}

	return nil
}

func (us UserService) SignupPlayer(ctx context.Context, input oapipublic.SignupPlayerRequestObject) (oapipublic.SignupPlayerResponseObject, error) {

	err := validateSignupPlayer(ctx, us.UserStore, input)
	if err != nil {
		return returnSignupPlayerError(err), nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Body.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	_, err = us.UserStore.CreateUser(ctx, nil, string(input.Body.Email), string(hash), string(models.UserTypePlayer))

	if err != nil {
		return nil, err
	}

	return oapipublic.SignupPlayer201Response{}, nil
}
