package models

import (
	"common/utils"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type Club struct {
	ID        string
	OwnerID   string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (c *Club) PostgresToModel(fromDb database.Club) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	ownerID, err := uuid.FromBytes(fromDb.OwnerID.Bytes[:])
	if err != nil {
		return err
	}

	c.ID = id.String()
	c.OwnerID = ownerID.String()
	c.Name = fromDb.Name
	c.CreatedAt = fromDb.CreatedAt.Time
	c.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}

func ClubFactory(count int, args map[string]interface{}) []Club {
	clubs := []Club{}

	ownerID, ok := args["ownerID"]
	if !ok {
		ownerID = uuid.NewString()
	}

	for i := 0; i < count; i++ {
		clubs = append(clubs, Club{
			OwnerID:   ownerID.(string),
			Name:      utils.NewString(10),
			CreatedAt: time.Now(),
		})
	}

	return clubs
}
