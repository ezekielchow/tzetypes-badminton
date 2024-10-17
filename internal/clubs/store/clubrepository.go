package clubs

import (
	"common/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ClubRepository interface {
	CreateClub(ctx context.Context, tx *pgx.Tx, toCreate models.Club) (models.Club, error)
	AddPlayerToClub(ctx context.Context, tx *pgx.Tx, playerID string, clubID string) error
	GetClubGivenOwnerId(ctx context.Context, tx *pgx.Tx, ownerID string) (models.Club, error)
	FindPlayerInClub(ctx context.Context, tx *pgx.Tx, clubID string, playerID string) (models.PlayerClub, error)
}
