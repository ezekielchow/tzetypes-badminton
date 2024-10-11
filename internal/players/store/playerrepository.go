package players

import (
	"common/models"
	"context"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, toCreate models.Player, passwordHash string) (models.Player, error)
	ListPlayers(ctx context.Context, ownerID *string, sortArrangement ListPlayersSort, offset int32, limit int32) ([]models.Player, int64, error)
}
