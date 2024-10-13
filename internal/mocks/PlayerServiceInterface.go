// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	oapiprivate "common/oapiprivate"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PlayerServiceInterface is an autogenerated mock type for the PlayerServiceInterface type
type PlayerServiceInterface struct {
	mock.Mock
}

// AddPlayer provides a mock function with given fields: ctx, input, ownerID
func (_m *PlayerServiceInterface) AddPlayer(ctx context.Context, input oapiprivate.AddPlayerRequestObject, ownerID string) (oapiprivate.AddPlayerResponseObject, error) {
	ret := _m.Called(ctx, input, ownerID)

	if len(ret) == 0 {
		panic("no return value specified for AddPlayer")
	}

	var r0 oapiprivate.AddPlayerResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.AddPlayerRequestObject, string) (oapiprivate.AddPlayerResponseObject, error)); ok {
		return rf(ctx, input, ownerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.AddPlayerRequestObject, string) oapiprivate.AddPlayerResponseObject); ok {
		r0 = rf(ctx, input, ownerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oapiprivate.AddPlayerResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapiprivate.AddPlayerRequestObject, string) error); ok {
		r1 = rf(ctx, input, ownerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlayers provides a mock function with given fields: ctx, input
func (_m *PlayerServiceInterface) ListPlayers(ctx context.Context, input oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for ListPlayers")
	}

	var r0 oapiprivate.ListPlayersResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.ListPlayersRequestObject) (oapiprivate.ListPlayersResponseObject, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.ListPlayersRequestObject) oapiprivate.ListPlayersResponseObject); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oapiprivate.ListPlayersResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapiprivate.ListPlayersRequestObject) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPlayerServiceInterface creates a new instance of PlayerServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlayerServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PlayerServiceInterface {
	mock := &PlayerServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
