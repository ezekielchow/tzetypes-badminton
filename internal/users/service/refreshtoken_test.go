package users

import (
	"common/middlewares"
	"common/models"
	"common/oapipublic"
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRefreshToken(t *testing.T) {
	ctx := context.Background()
	userService := InitService(ctx)

	t.Run("successfully refresh token", func(t *testing.T) {

		userID := uuid.New()
		session := models.Session{}
		sessionExpiresAt, err := session.NewSessionTokenExpiry()
		if err != nil {
			t.Fatalf("unable to get session token expiry: %s", err)
		}
		refreshExpiresAt, err := session.NewRefreshTokenExpiry()
		if err != nil {
			t.Fatalf("unable to get refresh token expiry: %s", err)
		}

		session, err = userService.SessionStore.CreateSession(ctx, nil, userID.String(), *sessionExpiresAt, *refreshExpiresAt)
		if err != nil {
			t.Fatalf("unable to create session: %s", err)
		}

		mockReq, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatalf("unable to create request: %s", err)
		}

		cookie := &http.Cookie{
			Name:     "refreshToken",
			Value:    session.RefreshToken,
			Expires:  session.RefreshTokenExpiresAt,
			HttpOnly: true,
			Secure:   false, // Set to true in production for HTTPS
			Path:     "/",
			SameSite: http.SameSiteNoneMode, // Prevent CSRF
		}

		// Add the cookie to the request
		mockReq.AddCookie(cookie)

		requestContext := context.Background()
		requestContext = context.WithValue(requestContext, middlewares.RequestKey, mockReq)

		res, err := userService.RefreshToken(requestContext, oapipublic.RefreshTokenRequestObject{})
		if err != nil {
			t.Fatalf("unable to refresh token: %s", err)
		}

		successRes, ok := res.(oapipublic.RefreshToken200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		updatedSession, err := userService.SessionStore.FindSessionToRefreshAccessToken(ctx, nil, session.RefreshToken)
		if err != nil {
			t.Fatalf("unable to get updated session: %s", err)
		}

		assert.Equal(t, updatedSession.SessionToken, successRes.SessionToken)
	})
}
