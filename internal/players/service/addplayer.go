package players

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

func (ps PlayerService) AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject, ownerID string) (oapiprivate.AddPlayerResponseObject, error) {
	// ownerID is logged in user

	tx, err := ps.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	user, err := ps.UserStore.CreateUser(ctx, &tx, "", "")
	if err != nil {
		return nil, err
	}

	player, err := ps.PlayerStore.CreatePlayer(ctx, &tx, models.Player{
		UserID: user.ID,
		Name:   input.Body.Name,
	}, "")

	if err != nil {
		return nil, err
	}

	club, err := ps.ClubStore.GetClubGivenOwnerId(ctx, &tx, ownerID)
	if err != nil {
		return nil, err
	}

	err = ps.ClubStore.AddPlayerToClub(ctx, &tx, player.ID, club.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.AddPlayer201JSONResponse{
		Id:        player.ID,
		Name:      player.Name,
		UserId:    player.UserID,
		CreatedAt: player.CreatedAt.String(),
		UpdatedAt: player.UpdatedAt.String(),
	}, nil
}
