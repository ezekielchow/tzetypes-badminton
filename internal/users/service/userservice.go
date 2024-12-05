package users

import (
	"common/models"
	"common/oapiprivate"
	"context"
	sessionstore "sessions/store"
	clubs "tzetypes-badminton/clubs/store"
	userstore "users/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServiceInterface interface {
	GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject, user models.User) (oapiprivate.GetLoggedInUserResponseObject, error)
}

type UserService struct {
	UserStore    userstore.UserRepository
	SessionStore sessionstore.SessionRepository
	ClubStore    clubs.ClubRepository
	PgxPool      *pgxpool.Pool
}
