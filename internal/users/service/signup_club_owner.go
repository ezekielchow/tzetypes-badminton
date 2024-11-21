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
	PasswordRepeatError = "repeated password has to be the same as password"
	EmailUsedError      = "email is already registered with another account"
)

func returnSignupError(err error) oapipublic.SignupClubOwnerdefaultJSONResponse {
	return oapipublic.SignupClubOwnerdefaultJSONResponse{
		Body:       oapipublic.Error{Message: err.Error()},
		StatusCode: http.StatusInternalServerError,
	}
}

func validateSignup(ctx context.Context, userStore userstore.UserRepository, input oapipublic.SignupClubOwnerRequestObject) error {

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

func (us UserService) SignupClubOwner(ctx context.Context, input oapipublic.SignupClubOwnerRequestObject) (oapipublic.SignupClubOwnerResponseObject, error) {

	tx, err := us.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	err = validateSignup(ctx, us.UserStore, input)
	if err != nil {
		return returnSignupError(err), nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Body.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	user, err := us.UserStore.CreateUser(ctx, &tx, string(input.Body.Email), string(hash), string(models.UserTypeClubOwner))

	if err != nil {
		return nil, err
	}

	_, err = us.ClubStore.CreateClub(ctx, &tx, models.Club{
		OwnerID: user.ID,
		Name:    user.Email,
	})

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapipublic.SignupClubOwner201Response{}, nil
}
