package userservice

import (
	"common/models"
	"common/oapipublic"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	LoginCredentialsInvalidError = "your email or password might be incorrect"
)

func returnLoginError(err error) oapipublic.LogindefaultJSONResponse {
	return oapipublic.LogindefaultJSONResponse{
		Body:       oapipublic.Error{Message: err.Error()},
		StatusCode: http.StatusInternalServerError,
	}
}

type Login200JSONResponse struct {
	SessionToken string `json:"session_token"`
	RefreshToken string `json:"refresh_token"`
}

func (response Login200JSONResponse) VisitLoginResponse(w http.ResponseWriter) error {
	// Set the cookie
	refreshTokenExpiry, err := models.Session{}.NewRefreshTokenExpiry()
	if err != nil {
		return err
	}

	isSecure, err := strconv.ParseBool(os.Getenv("IS_HTTPS"))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    response.RefreshToken,
		Expires:  *refreshTokenExpiry,
		HttpOnly: true,
		Secure:   isSecure, // Set to true in production for HTTPS
		Path:     "/",
		SameSite: http.SameSiteStrictMode, // Prevent CSRF
	})

	// Set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and write the JSON response body
	return json.NewEncoder(w).Encode(response)
}

func (us UserService) Login(ctx context.Context, input oapipublic.LoginRequestObject) (oapipublic.LoginResponseObject, error) {
	user, err := us.UserStore.FindUserWithEmail(ctx, nil, string(input.Body.Email))
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	if user.ID == "" {
		err = errors.New(LoginCredentialsInvalidError)
		return returnLoginError(err), nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Body.Password))
	if err != nil {
		// TODO: add log for all needed
		err = errors.New(LoginCredentialsInvalidError)
		return returnLoginError(err), nil
	}

	emptySession := models.Session{}

	sessionExpiresAt, err := emptySession.NewSessionTokenExpiry()
	if err != nil {
		return nil, err
	}
	refreshExpiresAt, err := emptySession.NewRefreshTokenExpiry()
	if err != nil {
		return nil, err
	}

	session, err := us.SessionStore.CreateSession(ctx, nil, user.ID, *sessionExpiresAt, *refreshExpiresAt)
	if err != nil {
		return nil, err
	}

	return Login200JSONResponse{
		SessionToken: session.SessionToken,
		RefreshToken: session.RefreshToken,
	}, nil
}
