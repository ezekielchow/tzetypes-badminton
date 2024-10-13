package common

import (
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	"fmt"
	player "players/service"
	userservice "users/service"
)

type CommonService struct {
	UserService   userservice.UserServiceInterface
	PlayerService player.PlayerServiceInterface
}

type Controller struct {
	Services CommonService
}

func NewController(cs CommonService) *Controller {
	return &Controller{Services: cs}
}

func (c Controller) Login(ctx context.Context, input oapipublic.LoginRequestObject) (oapipublic.LoginResponseObject, error) {
	return c.Services.UserService.Login(ctx, input)
}

func (c Controller) Signup(ctx context.Context, input oapipublic.SignupRequestObject) (oapipublic.SignupResponseObject, error) {
	return c.Services.UserService.Signup(ctx, input)
}

func (c Controller) RefreshToken(ctx context.Context, input oapipublic.RefreshTokenRequestObject) (oapipublic.RefreshTokenResponseObject, error) {
	return c.Services.UserService.RefreshToken(ctx, input)
}

func (c Controller) Dashboard(ctx context.Context, input oapiprivate.DashboardRequestObject) (oapiprivate.DashboardResponseObject, error) {
	fmt.Println("user?", ctx.Value(ContextUser))

	return oapiprivate.Dashboard204Response{}, nil
}

func (c Controller) Logout(ctx context.Context, input oapiprivate.LogoutRequestObject) (oapiprivate.LogoutResponseObject, error) {
	session, ok := ctx.Value(ContextSession).(models.Session)
	if !ok {
		return nil, fmt.Errorf("unable to convert session context")
	}

	return c.Services.UserService.Logout(ctx, input, session)
}

func (c Controller) AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject) (oapiprivate.AddPlayerResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.PlayerService.AddPlayer(ctx, input, user.ID)
}

func (c Controller) GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject) (oapiprivate.GetLoggedInUserResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.UserService.GetLoggedInUser(ctx, input, user)
}

func (c Controller) ListPlayers(ctx context.Context, input oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error) {
	return c.Services.PlayerService.ListPlayers(ctx, input)
}
