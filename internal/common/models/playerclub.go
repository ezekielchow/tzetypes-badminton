package models

import (
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type PlayerClub struct {
	ID        string
	PlayerID  string
	ClubID    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (pc *PlayerClub) PostgresToModel(fromDb database.PlayerClub) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	playerID, err := uuid.FromBytes(fromDb.PlayerID.Bytes[:])
	if err != nil {
		return err
	}

	clubID, err := uuid.FromBytes(fromDb.ClubID.Bytes[:])
	if err != nil {
		return err
	}

	pc.ID = id.String()
	pc.PlayerID = playerID.String()
	pc.ClubID = clubID.String()
	pc.CreatedAt = fromDb.CreatedAt.Time
	pc.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}
