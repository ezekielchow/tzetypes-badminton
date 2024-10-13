package clubs

import (
	"common/models"
	"common/utils"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ClubPostgres struct {
	Queries *database.Queries
}

func (cp ClubPostgres) CreateClub(ctx context.Context, tx *pgx.Tx, toCreate models.Club) (models.Club, error) {

	pgOwnerID := pgtype.UUID{}
	err := pgOwnerID.Scan(toCreate.OwnerID)
	if err != nil {
		return models.Club{}, err
	}

	created, err := cp.Queries.CreateClub(ctx, database.CreateClubParams{
		OwnerID: pgOwnerID,
		Name:    toCreate.Name,
	})

	if err != nil {
		return models.Club{}, err
	}

	club := models.Club{}
	err = club.PostgresToModel(created)
	if err != nil {
		return models.Club{}, err
	}

	return club, nil
}

func (cp ClubPostgres) AddPlayerToClub(ctx context.Context, tx *pgx.Tx, playerID string, clubID string) error {
	pgPlayerID, err := utils.StringToPgId(playerID)
	if err != nil {
		return err
	}

	pgClubID, err := utils.StringToPgId(clubID)
	if err != nil {
		return err
	}

	err = cp.Queries.AddPlayerToClub(ctx, database.AddPlayerToClubParams{
		PlayerID: pgPlayerID,
		ClubID:   pgClubID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (cp ClubPostgres) GetClubGivenOwnerId(ctx context.Context, tx *pgx.Tx, ownerID string) (models.Club, error) {
	pgOwnerID, err := utils.StringToPgId(ownerID)
	if err != nil {
		return models.Club{}, err
	}

	pgClub, err := cp.Queries.GetClubGivenOwnerId(ctx, pgOwnerID)
	if err != nil {
		return models.Club{}, err
	}

	club := models.Club{}
	err = club.PostgresToModel(pgClub)
	if err != nil {
		return models.Club{}, err
	}

	return club, nil
}
