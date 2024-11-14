// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ServerInterface is an autogenerated mock type for the ServerInterface type
type ServerInterface struct {
	mock.Mock
}

// GetGame provides a mock function with given fields: w, r, gameId
func (_m *ServerInterface) GetGame(w http.ResponseWriter, r *http.Request, gameId string) {
	_m.Called(w, r, gameId)
}

// Login provides a mock function with given fields: w, r
func (_m *ServerInterface) Login(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// RefreshToken provides a mock function with given fields: w, r
func (_m *ServerInterface) RefreshToken(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Signup provides a mock function with given fields: w, r
func (_m *ServerInterface) Signup(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewServerInterface creates a new instance of ServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerInterface {
	mock := &ServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
