package players

import (
	"common/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, tx *pgx.Tx, toCreate models.Player, passwordHash string) (models.Player, error)
	ListPlayers(ctx context.Context, tx *pgx.Tx, ownerID *string, sortArrangement ListPlayersSort, offset int32, limit int32) ([]models.Player, int64, error)
	FindUserWithName(ctx context.Context, tx *pgx.Tx, name string) (models.Player, error)
	AllPlayers(ctx context.Context, tx *pgx.Tx) ([]models.Player, error)
}
