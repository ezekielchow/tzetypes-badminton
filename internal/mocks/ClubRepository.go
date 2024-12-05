// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	models "common/models"
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// ClubRepository is an autogenerated mock type for the ClubRepository type
type ClubRepository struct {
	mock.Mock
}

// AddPlayerToClub provides a mock function with given fields: ctx, tx, playerID, clubID
func (_m *ClubRepository) AddPlayerToClub(ctx context.Context, tx *pgx.Tx, playerID string, clubID string) error {
	ret := _m.Called(ctx, tx, playerID, clubID)

	if len(ret) == 0 {
		panic("no return value specified for AddPlayerToClub")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, string) error); ok {
		r0 = rf(ctx, tx, playerID, clubID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateClub provides a mock function with given fields: ctx, tx, toCreate
func (_m *ClubRepository) CreateClub(ctx context.Context, tx *pgx.Tx, toCreate models.Club) (models.Club, error) {
	ret := _m.Called(ctx, tx, toCreate)

	if len(ret) == 0 {
		panic("no return value specified for CreateClub")
	}

	var r0 models.Club
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Club) (models.Club, error)); ok {
		return rf(ctx, tx, toCreate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, models.Club) models.Club); ok {
		r0 = rf(ctx, tx, toCreate)
	} else {
		r0 = ret.Get(0).(models.Club)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, models.Club) error); ok {
		r1 = rf(ctx, tx, toCreate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPlayerInClub provides a mock function with given fields: ctx, tx, clubID, playerID
func (_m *ClubRepository) FindPlayerInClub(ctx context.Context, tx *pgx.Tx, clubID string, playerID string) (models.PlayerClub, error) {
	ret := _m.Called(ctx, tx, clubID, playerID)

	if len(ret) == 0 {
		panic("no return value specified for FindPlayerInClub")
	}

	var r0 models.PlayerClub
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, string) (models.PlayerClub, error)); ok {
		return rf(ctx, tx, clubID, playerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string, string) models.PlayerClub); ok {
		r0 = rf(ctx, tx, clubID, playerID)
	} else {
		r0 = ret.Get(0).(models.PlayerClub)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string, string) error); ok {
		r1 = rf(ctx, tx, clubID, playerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClubGivenOwnerID provides a mock function with given fields: ctx, tx, ownerID
func (_m *ClubRepository) GetClubGivenOwnerID(ctx context.Context, tx *pgx.Tx, ownerID string) (models.Club, error) {
	ret := _m.Called(ctx, tx, ownerID)

	if len(ret) == 0 {
		panic("no return value specified for GetClubGivenOwnerID")
	}

	var r0 models.Club
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) (models.Club, error)); ok {
		return rf(ctx, tx, ownerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Tx, string) models.Club); ok {
		r0 = rf(ctx, tx, ownerID)
	} else {
		r0 = ret.Get(0).(models.Club)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgx.Tx, string) error); ok {
		r1 = rf(ctx, tx, ownerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClubRepository creates a new instance of ClubRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClubRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClubRepository {
	mock := &ClubRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
