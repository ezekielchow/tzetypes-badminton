package clubs

import (
	"common/models"
	"context"
	clubstoregenerated "tzetypes-badminton/clubs/store/generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type ClubPostgres struct {
	Queries *clubstoregenerated.Queries
}

func (cp ClubPostgres) CreateClub(ctx context.Context, toCreate models.Club) (models.Club, error) {

	pgOwnerID := pgtype.UUID{}
	err := pgOwnerID.Scan(toCreate.OwnerID)
	if err != nil {
		return models.Club{}, err
	}

	created, err := cp.Queries.CreateClub(ctx, clubstoregenerated.CreateClubParams{
		OwnerID: pgOwnerID,
		Name:    toCreate.Name,
	})

	if err != nil {
		return models.Club{}, err
	}

	club := models.Club{}
	club.PostgresToModel(created)

	return club, nil
}
