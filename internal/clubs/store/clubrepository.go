package clubs

import (
	"common/models"
	"context"
)

type ClubRepository interface {
	CreateClub(ctx context.Context, toCreate models.Club) (models.Club, error)
}
