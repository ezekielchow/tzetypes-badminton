package userservice

import (
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	sessionstore "sessions/store"
	userstore "users/store"
)

type UserServiceInterface interface {
	Signup(ctx context.Context, input oapipublic.SignupRequestObject) (oapipublic.SignupResponseObject, error)
	Login(ctx context.Context, input oapipublic.LoginRequestObject) (oapipublic.LoginResponseObject, error)
	RefreshToken(ctx context.Context, input oapipublic.RefreshTokenRequestObject) (oapipublic.RefreshTokenResponseObject, error)
	Logout(ctx context.Context, input oapiprivate.LogoutRequestObject, session models.Session) (oapiprivate.LogoutResponseObject, error)
	GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject, user models.User) (oapiprivate.GetLoggedInUserResponseObject, error)
}

type UserService struct {
	UserStore    userstore.UserRepository
	SessionStore sessionstore.SessionRepository
}
