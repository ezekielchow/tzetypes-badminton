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

// CreateOrUpdateGameHistory provides a mock function with given fields: ctx, tx, toCreate
func (_m *GameRepository) CreateOrUpdateGameHistory(ctx context.Context, tx *pgx.Tx, toCreate models.GameHistory) (models.GameHistory, error) {
	ret := _m.Called(ctx, tx, toCreate)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateGameHistory")
	}

	var r0 models.GameHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.GameHistory) (models.GameHistory, error)); ok {
		return rf(ctx, tx, toCreate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.GameHistory) models.GameHistory); ok {
		r0 = rf(ctx, tx, toCreate)
	} else {
		r0 = ret.Get(0).(models.GameHistory)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, models.GameHistory) error); ok {
		r1 = rf(ctx, tx, toCreate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStatistic provides a mock function with given fields: ctx, tx, gameID, toCreate
func (_m *GameRepository) CreateStatistic(ctx context.Context, tx *pgx.Tx, gameID string, toCreate models.GameStatistic) (models.GameStatistic, error) {
	ret := _m.Called(ctx, tx, gameID, toCreate)

	if len(ret) == 0 {
		panic("no return value specified for CreateStatistic")
	}

	var r0 models.GameStatistic
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, models.GameStatistic) (models.GameStatistic, error)); ok {
		return rf(ctx, tx, gameID, toCreate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, models.GameStatistic) models.GameStatistic); ok {
		r0 = rf(ctx, tx, gameID, toCreate)
	} else {
		r0 = ret.Get(0).(models.GameStatistic)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string, models.GameStatistic) error); ok {
		r1 = rf(ctx, tx, gameID, toCreate)
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

// EndGame provides a mock function with given fields: ctx, tx, id, isEnded
func (_m *GameRepository) EndGame(ctx context.Context, tx *pgx.Tx, id string, isEnded bool) error {
	ret := _m.Called(ctx, tx, id, isEnded)

	if len(ret) == 0 {
		panic("no return value specified for EndGame")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, bool) error); ok {
		r0 = rf(ctx, tx, id, isEnded)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGame provides a mock function with given fields: ctx, tx, id
func (_m *GameRepository) GetGame(ctx context.Context, tx *pgx.Tx, id string) (models.Game, error) {
	ret := _m.Called(ctx, tx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetGame")
	}

	var r0 models.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) (models.Game, error)); ok {
		return rf(ctx, tx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) models.Game); ok {
		r0 = rf(ctx, tx, id)
	} else {
		r0 = ret.Get(0).(models.Game)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string) error); ok {
		r1 = rf(ctx, tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameHistoryGivenUserIdAndGameId provides a mock function with given fields: ctx, tx, userID, gameID
func (_m *GameRepository) GetGameHistoryGivenUserIdAndGameId(ctx context.Context, tx *pgx.Tx, userID string, gameID string) (models.GameHistory, error) {
	ret := _m.Called(ctx, tx, userID, gameID)

	if len(ret) == 0 {
		panic("no return value specified for GetGameHistoryGivenUserIdAndGameId")
	}

	var r0 models.GameHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, string) (models.GameHistory, error)); ok {
		return rf(ctx, tx, userID, gameID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, string) models.GameHistory); ok {
		r0 = rf(ctx, tx, userID, gameID)
	} else {
		r0 = ret.Get(0).(models.GameHistory)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string, string) error); ok {
		r1 = rf(ctx, tx, userID, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameSteps provides a mock function with given fields: ctx, tx, gameID
func (_m *GameRepository) GetGameSteps(ctx context.Context, tx *pgx.Tx, gameID string) ([]models.GameStep, error) {
	ret := _m.Called(ctx, tx, gameID)

	if len(ret) == 0 {
		panic("no return value specified for GetGameSteps")
	}

	var r0 []models.GameStep
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) ([]models.GameStep, error)); ok {
		return rf(ctx, tx, gameID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) []models.GameStep); ok {
		r0 = rf(ctx, tx, gameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.GameStep)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string) error); ok {
		r1 = rf(ctx, tx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatisticsWithGameId provides a mock function with given fields: ctx, tx, gameID
func (_m *GameRepository) GetStatisticsWithGameId(ctx context.Context, tx *pgx.Tx, gameID string) (models.GameStatistic, error) {
	ret := _m.Called(ctx, tx, gameID)

	if len(ret) == 0 {
		panic("no return value specified for GetStatisticsWithGameId")
	}

	var r0 models.GameStatistic
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) (models.GameStatistic, error)); ok {
		return rf(ctx, tx, gameID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) models.GameStatistic); ok {
		r0 = rf(ctx, tx, gameID)
	} else {
		r0 = ret.Get(0).(models.GameStatistic)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string) error); ok {
		r1 = rf(ctx, tx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
