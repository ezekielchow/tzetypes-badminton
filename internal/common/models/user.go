package models

import (
	"common/utils"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func (u *User) PostgresToModel(fromDb database.User) error {
	uuid, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	u.ID = uuid.String()
	u.Email = fromDb.Email
	u.PasswordHash = *fromDb.PasswordHash
	u.CreatedAt = fromDb.CreatedAt.Time
	u.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}

func (u *User) Mock() {
	u.ID = uuid.NewString()
	u.Email = utils.NewEmail(6)
	u.PasswordHash = utils.NewString(24)
	u.CreatedAt = time.Now()
	u.UpdatedAt = nil
}
