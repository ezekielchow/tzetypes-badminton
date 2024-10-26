// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// DeleteGameStepsResponseObject is an autogenerated mock type for the DeleteGameStepsResponseObject type
type DeleteGameStepsResponseObject struct {
	mock.Mock
}

// VisitDeleteGameStepsResponse provides a mock function with given fields: w
func (_m *DeleteGameStepsResponseObject) VisitDeleteGameStepsResponse(w http.ResponseWriter) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for VisitDeleteGameStepsResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDeleteGameStepsResponseObject creates a new instance of DeleteGameStepsResponseObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeleteGameStepsResponseObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeleteGameStepsResponseObject {
	mock := &DeleteGameStepsResponseObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
