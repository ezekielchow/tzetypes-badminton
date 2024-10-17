package players

import (
	"common/models"
	"common/oapiprivate"
	"common/utils"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createRandomClubs(t *testing.T, ctx context.Context, playerService PlayerService, playerCount int) models.Club {
	owner, err := playerService.UserStore.CreateUser(ctx, nil, utils.NewEmail(10), "")
	if err != nil {
		t.Fatalf("unable to create owner user: %s", err)
	}

	club, err := playerService.ClubStore.CreateClub(ctx, nil, models.Club{
		OwnerID: owner.ID,
		Name:    utils.NewString(10),
	})
	if err != nil {
		t.Fatalf("unable to create club: %s", err)
	}

	for i := 0; i < playerCount; i++ {
		player, err := playerService.PlayerStore.CreatePlayer(ctx, nil, models.Player{
			UserID: uuid.New().String(),
			Name:   utils.NewString(10),
		}, "")
		if err != nil {
			t.Fatalf("unable to create player: %s", err)
		}

		err = playerService.ClubStore.AddPlayerToClub(ctx, nil, player.ID, club.ID)
		if err != nil {
			t.Fatalf("unable to add player to club: %s", err)
		}
	}

	return club
}

func TestListPlayers(t *testing.T) {
	ctx := context.Background()
	playerService := InitService(ctx)

	_ = createRandomClubs(t, ctx, playerService, 10)

	perPage := 10
	owner, err := playerService.UserStore.CreateUser(ctx, nil, utils.NewEmail(10), "")
	if err != nil {
		t.Fatalf("unable to create owner user: %s", err)
	}

	club, err := playerService.ClubStore.CreateClub(ctx, nil, models.Club{
		OwnerID: owner.ID,
		Name:    utils.NewString(10),
	})
	if err != nil {
		t.Fatalf("unable to create club: %s", err)
	}

	playerNames := [12]string{}

	t.Run("successfully list players no pagination", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			name := utils.NewString(10)
			playerNames[i] = name
			player, err := playerService.PlayerStore.CreatePlayer(ctx, nil, models.Player{
				UserID: uuid.New().String(),
				Name:   name,
			}, "")
			if err != nil {
				t.Fatalf("unable to create player: %s", err)
			}

			err = playerService.ClubStore.AddPlayerToClub(ctx, nil, player.ID, club.ID)
			if err != nil {
				t.Fatalf("unable to add player to club: %s", err)
			}
		}

		sort := "name_asc"
		res, err := playerService.ListPlayers(ctx, oapiprivate.ListPlayersRequestObject{
			Params: oapiprivate.ListPlayersParams{
				OwnerId:         &owner.ID,
				Page:            1,
				PageSize:        perPage,
				SortArrangement: &sort,
			},
		})
		if err != nil {
			t.Fatalf("unable to listplayers: %s", err)
		}
		successRes, ok := res.(oapiprivate.ListPlayers200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, 5, successRes.Pagination.TotalItems)
		assert.Equal(t, 1, successRes.Pagination.CurrentPage)
		assert.Equal(t, 10, successRes.Pagination.PageSize)
		assert.Equal(t, 1, successRes.Pagination.TotalPages)

		for _, p := range *successRes.Players {
			isExist := false

			for _, name := range playerNames {
				if name == p.Name {
					isExist = true
				}
			}
			if !isExist {
				t.Fatal("Listed wrong player")
			}
		}
	})

	t.Run("successfully list players with pagination", func(t *testing.T) {
		for i := 5; i < 12; i++ {
			name := utils.NewString(10)
			playerNames[i] = name
			player, err := playerService.PlayerStore.CreatePlayer(ctx, nil, models.Player{
				UserID: uuid.New().String(),
				Name:   name,
			}, "")
			if err != nil {
				t.Fatalf("unable to create player: %s", err)
			}

			err = playerService.ClubStore.AddPlayerToClub(ctx, nil, player.ID, club.ID)
			if err != nil {
				t.Fatalf("unable to add player to club: %s", err)
			}
		}

		sort := "name_asc"
		res, err := playerService.ListPlayers(ctx, oapiprivate.ListPlayersRequestObject{
			Params: oapiprivate.ListPlayersParams{
				OwnerId:         &owner.ID,
				Page:            1,
				PageSize:        perPage,
				SortArrangement: &sort,
			},
		})
		if err != nil {
			t.Fatalf("unable to listplayers: %s", err)
		}
		successRes, ok := res.(oapiprivate.ListPlayers200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, 12, successRes.Pagination.TotalItems)
		assert.Equal(t, 1, successRes.Pagination.CurrentPage)
		assert.Equal(t, 10, successRes.Pagination.PageSize)
		assert.Equal(t, 2, successRes.Pagination.TotalPages)
		assert.Equal(t, 10, len(*successRes.Players))

		for _, p := range *successRes.Players {
			isExist := false

			for _, name := range playerNames {
				if name == p.Name {
					isExist = true
				}
			}
			if !isExist {
				t.Fatal("Listed wrong player")
			}
		}

		res, err = playerService.ListPlayers(ctx, oapiprivate.ListPlayersRequestObject{
			Params: oapiprivate.ListPlayersParams{
				OwnerId:         &owner.ID,
				Page:            2,
				PageSize:        perPage,
				SortArrangement: &sort,
			},
		})
		if err != nil {
			t.Fatalf("unable to listplayers: %s", err)
		}
		successRes, ok = res.(oapiprivate.ListPlayers200JSONResponse)
		if !ok {
			t.Fatal("unable to convert response")
		}

		assert.Equal(t, 12, successRes.Pagination.TotalItems)
		assert.Equal(t, 2, successRes.Pagination.CurrentPage)
		assert.Equal(t, 10, successRes.Pagination.PageSize)
		assert.Equal(t, 2, successRes.Pagination.TotalPages)
		assert.Equal(t, 2, len(*successRes.Players))

		for _, p := range *successRes.Players {
			isExist := false

			for _, name := range playerNames {
				if name == p.Name {
					isExist = true
				}
			}
			if !isExist {
				t.Fatal("Listed wrong player")
			}
		}
	})
}
