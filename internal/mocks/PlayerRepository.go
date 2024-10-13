// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	models "common/models"
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"

	players "players/store"
)

// PlayerRepository is an autogenerated mock type for the PlayerRepository type
type PlayerRepository struct {
	mock.Mock
}

// CreatePlayer provides a mock function with given fields: ctx, tx, toCreate, passwordHash
func (_m *PlayerRepository) CreatePlayer(ctx context.Context, tx *pgx.Tx, toCreate models.Player, passwordHash string) (models.Player, error) {
	ret := _m.Called(ctx, tx, toCreate, passwordHash)

	if len(ret) == 0 {
		panic("no return value specified for CreatePlayer")
	}

	var r0 models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Player, string) (models.Player, error)); ok {
		return rf(ctx, tx, toCreate, passwordHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Player, string) models.Player); ok {
		r0 = rf(ctx, tx, toCreate, passwordHash)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, models.Player, string) error); ok {
		r1 = rf(ctx, tx, toCreate, passwordHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlayers provides a mock function with given fields: ctx, tx, ownerID, sortArrangement, offset, limit
func (_m *PlayerRepository) ListPlayers(ctx context.Context, tx *pgx.Tx, ownerID *string, sortArrangement players.ListPlayersSort, offset int32, limit int32) ([]models.Player, int64, error) {
	ret := _m.Called(ctx, tx, ownerID, sortArrangement, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ListPlayers")
	}

	var r0 []models.Player
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, *string, players.ListPlayersSort, int32, int32) ([]models.Player, int64, error)); ok {
		return rf(ctx, tx, ownerID, sortArrangement, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, *string, players.ListPlayersSort, int32, int32) []models.Player); ok {
		r0 = rf(ctx, tx, ownerID, sortArrangement, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, *string, players.ListPlayersSort, int32, int32) int64); ok {
		r1 = rf(ctx, tx, ownerID, sortArrangement, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *pgx.Tx, *string, players.ListPlayersSort, int32, int32) error); ok {
		r2 = rf(ctx, tx, ownerID, sortArrangement, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewPlayerRepository creates a new instance of PlayerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlayerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PlayerRepository {
	mock := &PlayerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
