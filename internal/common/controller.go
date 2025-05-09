package common

import (
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	"fmt"
	games "games/service"
	players "players/service"
	users "users/service"
)

type CommonService struct {
	UserService   users.UserServiceInterface
	PlayerService players.PlayerServiceInterface
	GameService   games.GameServiceInterface
}

type Controller struct {
	Services CommonService
}

func NewController(cs CommonService) *Controller {
	return &Controller{Services: cs}
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

func (c Controller) UpdatePlayerWithId(ctx context.Context, input oapiprivate.UpdatePlayerWithIdRequestObject) (oapiprivate.UpdatePlayerWithIdResponseObject, error) {
	return c.Services.PlayerService.UpdatePlayer(ctx, input)
}

func (c Controller) GetPlayerWithId(ctx context.Context, input oapiprivate.GetPlayerWithIdRequestObject) (oapiprivate.GetPlayerWithIdResponseObject, error) {
	return c.Services.PlayerService.GetPlayerWithId(ctx, input)
}

func (c Controller) StartGame(ctx context.Context, input oapiprivate.StartGameRequestObject) (oapiprivate.StartGameResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.GameService.StartGame(ctx, input, user)
}

func (c Controller) AddGameSteps(ctx context.Context, input oapiprivate.AddGameStepsRequestObject) (oapiprivate.AddGameStepsResponseObject, error) {
	return c.Services.GameService.AddGameSteps(ctx, input)
}

func (c Controller) DeleteGameSteps(ctx context.Context, input oapiprivate.DeleteGameStepsRequestObject) (oapiprivate.DeleteGameStepsResponseObject, error) {
	return c.Services.GameService.DeleteGameSteps(ctx, input)
}

func (c Controller) EndGame(ctx context.Context, input oapiprivate.EndGameRequestObject) (oapiprivate.EndGameResponseObject, error) {
	return c.Services.GameService.EndGame(ctx, input)
}

func (c Controller) GetGame(ctx context.Context, input oapiprivate.GetGameRequestObject) (oapiprivate.GetGameResponseObject, error) {
	return c.Services.GameService.GetGame(ctx, input)
}

func (c Controller) CreateOrUpdateGameHistory(ctx context.Context, input oapiprivate.CreateOrUpdateGameHistoryRequestObject) (oapiprivate.CreateOrUpdateGameHistoryResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.GameService.CreateOrUpdateGameHistory(ctx, input, user)
}

func (c Controller) GetGameHistory(ctx context.Context, input oapiprivate.GetGameHistoryRequestObject) (oapiprivate.GetGameHistoryResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.GameService.GetGameHistory(ctx, input, user)
}

func (c Controller) GenerateRecentStatistics(ctx context.Context, input oapipublic.GenerateRecentStatisticsRequestObject) (oapipublic.GenerateRecentStatisticsResponseObject, error) {
	return c.Services.GameService.GenerateRecentStatistics(ctx, input)
}

func (c Controller) GetRecentStatistics(ctx context.Context, input oapiprivate.GetRecentStatisticsRequestObject) (oapiprivate.GetRecentStatisticsResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.GameService.GetRecentStatistics(ctx, input, user)
}

func (c Controller) EndAbandonedGames(ctx context.Context, input oapipublic.EndAbandonedGamesRequestObject) (oapipublic.EndAbandonedGamesResponseObject, error) {
	return c.Services.GameService.EndAbandonedGames(ctx, input)
}

func (c Controller) ListActiveGames(ctx context.Context, input oapiprivate.ListActiveGamesRequestObject) (oapiprivate.ListActiveGamesResponseObject, error) {
	user, ok := ctx.Value(ContextUser).(models.User)
	if !ok {
		return nil, fmt.Errorf("unable to convert user context")
	}

	return c.Services.GameService.ListActiveGames(ctx, input, user)
}

func (c Controller) GetGameStatistics(ctx context.Context, input oapipublic.GetGameStatisticsRequestObject) (oapipublic.GetGameStatisticsResponseObject, error) {
	return c.Services.GameService.GetGameStatistics(ctx, input)
}

func (c Controller) UpdateInstagramFeed(ctx context.Context, input oapipublic.UpdateInstagramFeedRequestObject) (oapipublic.UpdateInstagramFeedResponseObject, error) {
	return c.Services.GameService.UpdateInstagramFeed(ctx, input)
}

func (c Controller) GetInstagramFeed(ctx context.Context, input oapipublic.GetInstagramFeedRequestObject) (oapipublic.GetInstagramFeedResponseObject, error) {
	return c.Services.GameService.GetInstagramFeed(ctx, input)
}
