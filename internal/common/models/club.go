package models

import (
	"time"
	clubstoregenerated "tzetypes-badminton/clubs/store/generated"

	"github.com/google/uuid"
)

type Club struct {
	ID        string
	OwnerID   string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (c *Club) PostgresToModel(fromDb clubstoregenerated.Club) error {
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
