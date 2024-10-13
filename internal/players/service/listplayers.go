package player

import (
	"common/oapiprivate"
	"context"
	players "players/store"
)

func (ps PlayerService) ListPlayers(ctx context.Context, input oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error) {
	offset := (input.Params.Page - 1) * input.Params.PageSize

	players, totalCount, err := ps.PlayerStore.ListPlayers(ctx, nil, input.Params.OwnerId, players.ListPlayersSort(*input.Params.SortArrangement), int32(offset), int32(input.Params.PageSize))
	totalPages := totalCount / int64(input.Params.PageSize)

	if err != nil {
		return nil, err
	}

	apiPlayers := []oapiprivate.Player{}
	for _, p := range players {
		apiPlayers = append(apiPlayers, oapiprivate.Player{
			Id:        p.ID,
			Name:      p.Name,
			UserId:    p.UserID,
			CreatedAt: p.CreatedAt.String(),
			UpdatedAt: p.UpdatedAt.String(),
		})
	}

	return oapiprivate.ListPlayers200JSONResponse{
		Pagination: &struct {
			CurrentPage int `json:"currentPage"`
			PageSize    int `json:"pageSize"`
			TotalItems  int `json:"totalItems"`
			TotalPages  int `json:"totalPages"`
		}{
			CurrentPage: input.Params.Page,
			PageSize:    input.Params.PageSize,
			TotalItems:  int(totalCount),
			TotalPages:  int(totalPages),
		},
		Players: &apiPlayers,
	}, nil
}
