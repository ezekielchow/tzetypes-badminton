package players

import (
	"common/models"
	"context"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, toCreate models.Player, passwordHash string) (models.Player, error)
}
