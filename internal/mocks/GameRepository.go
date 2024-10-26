// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "common/models"

	pgx "github.com/jackc/pgx/v5"
)

// GameRepository is an autogenerated mock type for the GameRepository type
type GameRepository struct {
	mock.Mock
}

// CreateGame provides a mock function with given fields: ctx, tx, toCreate
func (_m *GameRepository) CreateGame(ctx context.Context, tx *pgx.Tx, toCreate models.Game) (models.Game, error) {
	ret := _m.Called(ctx, tx, toCreate)

	if len(ret) == 0 {
		panic("no return value specified for CreateGame")
	}

	var r0 models.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Game) (models.Game, error)); ok {
		return rf(ctx, tx, toCreate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Game) models.Game); ok {
		r0 = rf(ctx, tx, toCreate)
	} else {
		r0 = ret.Get(0).(models.Game)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, models.Game) error); ok {
		r1 = rf(ctx, tx, toCreate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGameStep provides a mock function with given fields: ctx, tx, toCreate
func (_m *GameRepository) CreateGameStep(ctx context.Context, tx *pgx.Tx, toCreate models.GameStep) (models.GameStep, error) {
	ret := _m.Called(ctx, tx, toCreate)

	if len(ret) == 0 {
		panic("no return value specified for CreateGameStep")
	}

	var r0 models.GameStep
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.GameStep) (models.GameStep, error)); ok {
		return rf(ctx, tx, toCreate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.GameStep) models.GameStep); ok {
		r0 = rf(ctx, tx, toCreate)
	} else {
		r0 = ret.Get(0).(models.GameStep)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, models.GameStep) error); ok {
		r1 = rf(ctx, tx, toCreate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGameStep provides a mock function with given fields: ctx, tx, id
func (_m *GameRepository) DeleteGameStep(ctx context.Context, tx *pgx.Tx, id string) error {
	ret := _m.Called(ctx, tx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGameStep")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) error); ok {
		r0 = rf(ctx, tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGameRepository creates a new instance of GameRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGameRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *GameRepository {
	mock := &GameRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
