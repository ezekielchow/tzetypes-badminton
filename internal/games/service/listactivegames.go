package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (gs GameService) ListActiveGames(ctx context.Context, input oapiprivate.ListActiveGamesRequestObject, user models.User) (oapiprivate.ListActiveGamesResponseObject, error) {
	club, err := gs.ClubStore.GetClubGivenOwnerID(ctx, nil, user.ID)
	if err != nil {
		return nil, err
	}

	games, err := gs.GameStore.GetActiveGames(ctx, nil, club.ID)
	if err != nil {
		return nil, err
	}

	apiGames := []oapiprivate.Game{}
	for _, game := range games {
		apiGames = append(apiGames, game.ModelToAPI())
	}

	return oapiprivate.ListActiveGames200JSONResponse{
		Games: &apiGames,
	}, nil
}
