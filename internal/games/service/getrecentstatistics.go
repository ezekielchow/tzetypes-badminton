package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (gs GameService) GetRecentStatistics(ctx context.Context, input oapiprivate.GetRecentStatisticsRequestObject, user models.User) (oapiprivate.GetRecentStatisticsResponseObject, error) {
	grs, err := gs.GameStore.GetGameRecentStatisticWithUserId(ctx, nil, user.ID)
	if err != nil {
		return nil, err
	}

	res := grs.ModelToAPI()
	return oapiprivate.GetRecentStatistics200JSONResponse{
		GetRecentStatisticsResponseSchemaJSONResponse: oapiprivate.GetRecentStatisticsResponseSchemaJSONResponse{
			GameRecentStatistics: res,
		},
	}, nil
}
