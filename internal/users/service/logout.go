package userservice

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Logout204Response struct {
}

func (response Logout204Response) VisitLogoutResponse(w http.ResponseWriter) error {
	isSecure, err := strconv.ParseBool(os.Getenv("IS_HTTPS"))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Expires:  time.Now().Add(0 - time.Minute*time.Duration(10)),
		HttpOnly: true,
		Secure:   isSecure, // Set to true in production for HTTPS
		Path:     "/",
		SameSite: http.SameSiteStrictMode, // Prevent CSRF
	})

	// Set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	// Encode and write the JSON response body
	return json.NewEncoder(w).Encode(response)
}

func (us UserService) Logout(ctx context.Context, input oapiprivate.LogoutRequestObject, session models.Session) (oapiprivate.LogoutResponseObject, error) {

	err := us.SessionStore.DeleteSession(ctx, session.ID)
	if err != nil {
		return nil, err
	}

	return Logout204Response{}, nil
}
