package players

import (
	"common/models"
	"context"
	playerstoregenerated "players/store/generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListPlayersSort string

const (
	ListPlayersSortNameAsc  ListPlayersSort = "name_asc"
	ListPlayersSortNameDesc ListPlayersSort = "name_desc"
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

func (pp PlayerPostgres) ListPlayers(ctx context.Context, ownerID *string, sortArrangement ListPlayersSort, offset int32, limit int32) ([]models.Player, int64, error) {

	var pgOwnerID pgtype.UUID

	if ownerID != nil {
		pgOwnerID = pgtype.UUID{}
		err := pgOwnerID.Scan(ownerID)
		if err != nil {
			return []models.Player{}, 0, err
		}
	}

	dbPlayers, err := pp.Queries.ListPlayers(ctx, playerstoregenerated.ListPlayersParams{
		OwnerID:         pgOwnerID,
		SortArrangement: string(sortArrangement),
		OffsetCount:     offset,
		LimitCount:      limit,
	})
	if err != nil {
		return []models.Player{}, 0, err
	}

	players := []models.Player{}
	for _, row := range dbPlayers {
		player := models.Player{}
		err := player.PostgresToModel(playerstoregenerated.Player{
			ID:        row.ID,
			UserID:    row.UserID,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
		if err != nil {
			return []models.Player{}, 0, err
		}

		players = append(players, player)
	}

	totalCount := int64(0)
	if len(dbPlayers) > 0 {
		totalCount = dbPlayers[0].TotalCount
	}

	return players, totalCount, nil
}
