// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	oapipublic "common/oapipublic"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// StrictServerInterface is an autogenerated mock type for the StrictServerInterface type
type StrictServerInterface struct {
	mock.Mock
}

// EndAbandonedGames provides a mock function with given fields: ctx, request
func (_m *StrictServerInterface) EndAbandonedGames(ctx context.Context, request oapipublic.EndAbandonedGamesRequestObject) (oapipublic.EndAbandonedGamesResponseObject, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for EndAbandonedGames")
	}

	var r0 oapipublic.EndAbandonedGamesResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.EndAbandonedGamesRequestObject) (oapipublic.EndAbandonedGamesResponseObject, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.EndAbandonedGamesRequestObject) oapipublic.EndAbandonedGamesResponseObject); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oapipublic.EndAbandonedGamesResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapipublic.EndAbandonedGamesRequestObject) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateRecentStatistics provides a mock function with given fields: ctx, request
func (_m *StrictServerInterface) GenerateRecentStatistics(ctx context.Context, request oapipublic.GenerateRecentStatisticsRequestObject) (oapipublic.GenerateRecentStatisticsResponseObject, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GenerateRecentStatistics")
	}

	var r0 oapipublic.GenerateRecentStatisticsResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.GenerateRecentStatisticsRequestObject) (oapipublic.GenerateRecentStatisticsResponseObject, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.GenerateRecentStatisticsRequestObject) oapipublic.GenerateRecentStatisticsResponseObject); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oapipublic.GenerateRecentStatisticsResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapipublic.GenerateRecentStatisticsRequestObject) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGame provides a mock function with given fields: ctx, request
func (_m *StrictServerInterface) GetGame(ctx context.Context, request oapipublic.GetGameRequestObject) (oapipublic.GetGameResponseObject, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetGame")
	}

	var r0 oapipublic.GetGameResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.GetGameRequestObject) (oapipublic.GetGameResponseObject, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapipublic.GetGameRequestObject) oapipublic.GetGameResponseObject); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oapipublic.GetGameResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapipublic.GetGameRequestObject) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStrictServerInterface creates a new instance of StrictServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStrictServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StrictServerInterface {
	mock := &StrictServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
