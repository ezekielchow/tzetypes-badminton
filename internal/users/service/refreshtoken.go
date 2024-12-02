package users

import (
	"common/middlewares"
	"common/models"
	"common/oapipublic"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	RefreshTokenExpiredError    = "refresh token has expired"
	NoMatchingSessionFoundError = "no matching session found"
)

func returnError(err error, statusCode int) oapipublic.RefreshTokendefaultJSONResponse {
	return oapipublic.RefreshTokendefaultJSONResponse{
		Body:       oapipublic.Error{Message: err.Error()},
		StatusCode: statusCode,
	}
}

func validate(session models.Session) error {

	if session.ID == "" {
		return errors.New(NoMatchingSessionFoundError)
	}

	if time.Now().After(session.RefreshTokenExpiresAt) {
		return errors.New(RefreshTokenExpiredError)
	}

	return nil
}

func (us UserService) RefreshToken(ctx context.Context, input oapipublic.RefreshTokenRequestObject) (oapipublic.RefreshTokenResponseObject, error) {

	tx, err := us.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	req, ok := ctx.Value(middlewares.RequestKey).(*http.Request)
	if !ok {
		return nil, fmt.Errorf("could not retrieve request from context")
	}

	refreshTokenCookie, err := req.Cookie("refreshToken")
	if err != nil {
		return nil, err
	}

	session, err := us.SessionStore.FindSessionToRefreshAccessToken(ctx, &tx, refreshTokenCookie.Value)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	if err = validate(session); err != nil {
		return returnError(err, http.StatusUnauthorized), nil
	}

	emptySession := models.Session{}
	newSessionExpiry, err := emptySession.NewSessionTokenExpiry()
	if err != nil {
		return nil, err
	}
	newSessionToken := uuid.NewString()
	_, err = us.SessionStore.UpdateSessionWithRefreshToken(ctx, &tx, refreshTokenCookie.Value, *newSessionExpiry, newSessionToken)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapipublic.RefreshToken200JSONResponse{
		SessionToken: newSessionToken,
	}, nil
}
