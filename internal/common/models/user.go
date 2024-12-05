package models

import (
	"common/utils"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type AccountTier string

const (
	AccountTierPlayer    AccountTier = "player"
	AccountTierSmallClub AccountTier = "small_club"
	AccountTierBigClub   AccountTier = "big_club"
)

type User struct {
	ID          string
	FirebaseUID string
	Email       string
	AccountTier string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (u *User) PostgresToModel(fromDb database.User) error {
	uuid, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	u.ID = uuid.String()
	u.FirebaseUID = fromDb.FirebaseUid
	u.Email = fromDb.Email
	u.AccountTier = fromDb.AccountTier
	u.CreatedAt = fromDb.CreatedAt.Time
	u.UpdatedAt = &fromDb.UpdatedAt.Time

	return nil
}

func (u *User) Mock() {
	u.ID = uuid.NewString()
	u.FirebaseUID = utils.NewString(10)
	u.Email = utils.NewEmail(6)
	u.AccountTier = utils.NewString(10)
	u.CreatedAt = time.Now()
	u.UpdatedAt = nil
}
