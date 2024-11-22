// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// CreateOrUpdateGameHistoryResponseObject is an autogenerated mock type for the CreateOrUpdateGameHistoryResponseObject type
type CreateOrUpdateGameHistoryResponseObject struct {
	mock.Mock
}

// VisitCreateOrUpdateGameHistoryResponse provides a mock function with given fields: w
func (_m *CreateOrUpdateGameHistoryResponseObject) VisitCreateOrUpdateGameHistoryResponse(w http.ResponseWriter) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for VisitCreateOrUpdateGameHistoryResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCreateOrUpdateGameHistoryResponseObject creates a new instance of CreateOrUpdateGameHistoryResponseObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateOrUpdateGameHistoryResponseObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateOrUpdateGameHistoryResponseObject {
	mock := &CreateOrUpdateGameHistoryResponseObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
