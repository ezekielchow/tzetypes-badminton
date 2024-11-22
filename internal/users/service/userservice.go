package users

import (
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	sessionstore "sessions/store"
	clubs "tzetypes-badminton/clubs/store"
	userstore "users/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServiceInterface interface {
	SignupClubOwner(ctx context.Context, input oapipublic.SignupClubOwnerRequestObject) (oapipublic.SignupClubOwnerResponseObject, error)
	Login(ctx context.Context, input oapipublic.LoginRequestObject) (oapipublic.LoginResponseObject, error)
	RefreshToken(ctx context.Context, input oapipublic.RefreshTokenRequestObject) (oapipublic.RefreshTokenResponseObject, error)
	Logout(ctx context.Context, input oapiprivate.LogoutRequestObject, session models.Session) (oapiprivate.LogoutResponseObject, error)
	GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject, user models.User) (oapiprivate.GetLoggedInUserResponseObject, error)
	SignupPlayer(ctx context.Context, input oapipublic.SignupPlayerRequestObject) (oapipublic.SignupPlayerResponseObject, error)
}

type UserService struct {
	UserStore    userstore.UserRepository
	SessionStore sessionstore.SessionRepository
	ClubStore    clubs.ClubRepository
	PgxPool      *pgxpool.Pool
}
