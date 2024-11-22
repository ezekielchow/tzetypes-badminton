package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"database/sql"
	"net/http"
	"strings"
)

const (
	GetGameHistoryNotFoundError = "unable to find player with given ID"
)

func (gs GameService) GetGameHistory(ctx context.Context, input oapiprivate.GetGameHistoryRequestObject, user models.User) (oapiprivate.GetGameHistoryResponseObject, error) {
	gameHistory, err := gs.GameStore.GetGameHistoryGivenUserIdAndGameId(ctx, nil, user.ID, input.GameId)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	if gameHistory.ID == "" {
		return oapiprivate.GetGameHistorydefaultJSONResponse{
			Body: oapiprivate.Error{
				Message: GetGameHistoryNotFoundError,
			},
			StatusCode: http.StatusNotFound,
		}, nil
	}

	apiRes := gameHistory.ModelToAPI()

	return oapiprivate.GetGameHistory200JSONResponse{
		GetGameHistoryResponseSchemaJSONResponse: oapiprivate.GetGameHistoryResponseSchemaJSONResponse{
			GameHistory: apiRes,
		},
	}, nil
}
