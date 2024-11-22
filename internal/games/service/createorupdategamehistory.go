package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (gs GameService) CreateOrUpdateGameHistory(ctx context.Context, input oapiprivate.CreateOrUpdateGameHistoryRequestObject, user models.User) (oapiprivate.CreateOrUpdateGameHistoryResponseObject, error) {
	gameHistory, err := gs.GameStore.CreateOrUpdateGameHistory(ctx, nil, models.GameHistory{
		UserID:         user.ID,
		GameID:         input.GameId,
		PlayerPosition: string(input.Body.PlayerPosition),
	})

	if err != nil {
		return nil, err
	}

	apiRes := gameHistory.ModelToAPI()

	return oapiprivate.CreateOrUpdateGameHistory200JSONResponse{
		CreateOrUpdateGameHistoryResponseSchemaJSONResponse: oapiprivate.CreateOrUpdateGameHistoryResponseSchemaJSONResponse{
			GameHistory: apiRes,
		},
	}, nil
}
