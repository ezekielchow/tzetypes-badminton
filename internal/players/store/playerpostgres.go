package players

import (
	"common/models"
	"context"
	playerstoregenerated "players/store/generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type PlayerPostgres struct {
	Queries *playerstoregenerated.Queries
}

func (pp PlayerPostgres) CreatePlayer(ctx context.Context, toCreate models.Player, passwordHash string) (models.Player, error) {

	pgUserID := pgtype.UUID{}
	err := pgUserID.Scan(toCreate.UserID)
	if err != nil {
		return models.Player{}, err
	}

	created, err := pp.Queries.CreatePlayer(ctx, playerstoregenerated.CreatePlayerParams{
		UserID: pgUserID,
		Name:   toCreate.Name,
	})

	if err != nil {
		return models.Player{}, err
	}

	player := models.Player{}
	player.PostgresToModel(created)

	return player, nil
}
