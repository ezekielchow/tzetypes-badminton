package players

import (
	"common/models"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListPlayersSort string

const (
	ListPlayersSortNameAsc  ListPlayersSort = "name_asc"
	ListPlayersSortNameDesc ListPlayersSort = "name_desc"
)

type PlayerPostgres struct {
	Queries *database.Queries
}

func (pp PlayerPostgres) CreatePlayer(ctx context.Context, tx *pgx.Tx, toCreate models.Player, passwordHash string) (models.Player, error) {

	queries := pp.Queries

	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID := pgtype.UUID{}
	err := pgUserID.Scan(toCreate.UserID)
	if err != nil {
		return models.Player{}, err
	}

	created, err := queries.CreatePlayer(ctx, database.CreatePlayerParams{
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

func (pp PlayerPostgres) ListPlayers(ctx context.Context, tx *pgx.Tx, ownerID *string, sortArrangement ListPlayersSort, offset int32, limit int32) ([]models.Player, int64, error) {

	queries := pp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	var pgOwnerID pgtype.UUID

	if ownerID != nil {
		pgOwnerID = pgtype.UUID{}
		err := pgOwnerID.Scan(ownerID)
		if err != nil {
			return []models.Player{}, 0, err
		}
	}

	dbPlayers, err := queries.ListPlayers(ctx, database.ListPlayersParams{
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
		err := player.PostgresToModel(database.Player{
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
