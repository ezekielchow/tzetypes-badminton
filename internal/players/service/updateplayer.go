package players

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (ps PlayerService) UpdatePlayer(ctx context.Context, input oapiprivate.UpdatePlayerWithIdRequestObject) (oapiprivate.UpdatePlayerWithIdResponseObject, error) {
	created, err := ps.PlayerStore.UpdatePlayer(ctx, nil, models.Player{
		ID:   input.Id,
		Name: input.Body.Name,
	})
	if err != nil {
		return nil, err
	}

	return oapiprivate.UpdatePlayerWithId200JSONResponse{
		Id:        created.ID,
		Name:      created.Name,
		UserId:    created.UserID,
		CreatedAt: created.CreatedAt.String(),
		UpdatedAt: created.UpdatedAt.String(),
	}, nil
}
