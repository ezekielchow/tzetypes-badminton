package users

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestLgout(t *testing.T) {
	ctx := context.Background()
	userService := InitService(ctx)

	t.Run("successful logout", func(t *testing.T) {

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

		res, err := userService.Logout(ctx, oapiprivate.LogoutRequestObject{}, session)
		if err != nil {
			t.Fatalf("unable to logout: %s", err)
		}
		_, ok := res.(Logout204Response)
		if !ok {
			t.Fatal("unable to convert response")
		}

		found, err := userService.SessionStore.FindSessionWithSessionID(ctx, nil, session.ID)
		if found.ID != "" {
			t.Fatalf("session not deleted: %s", err)
		}
	})
}
