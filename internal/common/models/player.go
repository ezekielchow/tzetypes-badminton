package models

import (
	playerstore "players/store/generated"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Player struct {
	ID        string
	UserID    string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (p *Player) PostgresToModel(fromDb playerstore.Player) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	userID, err := uuid.FromBytes(fromDb.UserID.Bytes[:])
	if err != nil {
		return err
	}

	p.ID = id.String()
	p.UserID = userID.String()
	p.Name = fromDb.Name
	p.CreatedAt = fromDb.CreatedAt.Time
	p.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}

func (p *Player) ModelToPostgres(model Player) (playerstore.Player, error) {

	dbPlayer := playerstore.Player{}

	id := pgtype.UUID{}
	err := id.Scan(model.ID)
	if err != nil {
		return dbPlayer, err
	}

	userID := pgtype.UUID{}
	err = userID.Scan(model.UserID)
	if err != nil {
		return dbPlayer, err
	}

	createdAt := pgtype.Timestamp{}
	err = createdAt.Scan(model.CreatedAt)
	if err != nil {
		return dbPlayer, err
	}

	updatedAt := pgtype.Timestamp{}
	err = updatedAt.Scan(model.UpdatedAt)
	if err != nil {
		return dbPlayer, err
	}

	dbPlayer.ID = id
	dbPlayer.UserID = userID
	dbPlayer.Name = model.Name
	dbPlayer.CreatedAt = createdAt
	dbPlayer.UpdatedAt = updatedAt

	return dbPlayer, nil
}
