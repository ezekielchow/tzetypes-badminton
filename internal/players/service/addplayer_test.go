package players

import (
	"common/models"
	"common/oapiprivate"
	"common/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPlayer(t *testing.T) {

	ctx := context.Background()
	playerService := InitService(ctx)
	t.Run("sucessfully add player", func(t *testing.T) {

		name := utils.NewString(10)
		ownerEmail := utils.NewEmail(10)

		user, err := playerService.UserStore.CreateUser(ctx, nil, ownerEmail, "")
		if err != nil {
			t.Fatalf("unable to create user:%s", err)
		}

		_, err = playerService.ClubStore.CreateClub(ctx, nil, models.Club{
			OwnerID: user.ID,
			Name:    user.Email,
		})
		if err != nil {
			t.Fatalf("unable to create owners club:%s", err)
		}

		res, err := playerService.AddPlayer(ctx, oapiprivate.AddPlayerRequestObject{
			Body: &oapiprivate.AddPlayerJSONRequestBody{
				Name: name,
			},
		}, user.ID)
		if err != nil {
			t.Fatalf("unable to add player:%s", err)
		}
		_, ok := res.(oapiprivate.AddPlayer201Response)
		if !ok {
			t.Fatal("unable to convert response")
		}

		player, err := playerService.PlayerStore.FindUserWithName(ctx, nil, name)
		if err != nil {
			t.Fatalf("unable to find player: %s", err)
		}
		assert.Equal(t, name, player.Name)

		_, err = playerService.UserStore.FindUserWithID(ctx, nil, player.UserID)
		if err != nil {
			t.Fatalf("unable to find user: %s", err)
		}

		club, err := playerService.ClubStore.GetClubGivenOwnerId(ctx, nil, user.ID)
		if err != nil {
			t.Fatalf("unable to find club: %s", err)
		}

		_, err = playerService.ClubStore.FindPlayerInClub(ctx, nil, club.ID, player.ID)
		if err != nil {
			t.Fatalf("player wasnt added to club: %s", err)
		}
	})
}
