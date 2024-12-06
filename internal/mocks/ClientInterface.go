// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	io "io"

	mock "github.com/stretchr/testify/mock"

	oapiprivate "common/oapiprivate"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// AddGameSteps provides a mock function with given fields: ctx, gameId, body, reqEditors
func (_m *ClientInterface) AddGameSteps(ctx context.Context, gameId string, body oapiprivate.AddGameStepsJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddGameSteps")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.AddGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.AddGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, oapiprivate.AddGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddGameStepsWithBody provides a mock function with given fields: ctx, gameId, contentType, body, reqEditors
func (_m *ClientInterface) AddGameStepsWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddGameStepsWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPlayer provides a mock function with given fields: ctx, body, reqEditors
func (_m *ClientInterface) AddPlayer(ctx context.Context, body oapiprivate.AddPlayerJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddPlayer")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.AddPlayerJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.AddPlayerJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapiprivate.AddPlayerJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPlayerWithBody provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *ClientInterface) AddPlayerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddPlayerWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateGameHistory provides a mock function with given fields: ctx, gameId, body, reqEditors
func (_m *ClientInterface) CreateOrUpdateGameHistory(ctx context.Context, gameId string, body oapiprivate.CreateOrUpdateGameHistoryJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateGameHistory")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.CreateOrUpdateGameHistoryJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.CreateOrUpdateGameHistoryJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, oapiprivate.CreateOrUpdateGameHistoryJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateGameHistoryWithBody provides a mock function with given fields: ctx, gameId, contentType, body, reqEditors
func (_m *ClientInterface) CreateOrUpdateGameHistoryWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateGameHistoryWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGameSteps provides a mock function with given fields: ctx, gameId, body, reqEditors
func (_m *ClientInterface) DeleteGameSteps(ctx context.Context, gameId string, body oapiprivate.DeleteGameStepsJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGameSteps")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.DeleteGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.DeleteGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, oapiprivate.DeleteGameStepsJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGameStepsWithBody provides a mock function with given fields: ctx, gameId, contentType, body, reqEditors
func (_m *ClientInterface) DeleteGameStepsWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGameStepsWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndGame provides a mock function with given fields: ctx, gameId, body, reqEditors
func (_m *ClientInterface) EndGame(ctx context.Context, gameId string, body oapiprivate.EndGameJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EndGame")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.EndGameJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.EndGameJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, oapiprivate.EndGameJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndGameWithBody provides a mock function with given fields: ctx, gameId, contentType, body, reqEditors
func (_m *ClientInterface) EndGameWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EndGameWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameHistory provides a mock function with given fields: ctx, gameId, reqEditors
func (_m *ClientInterface) GetGameHistory(ctx context.Context, gameId string, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gameId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGameHistory")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, gameId, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, gameId, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, gameId, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoggedInUser provides a mock function with given fields: ctx, reqEditors
func (_m *ClientInterface) GetLoggedInUser(ctx context.Context, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLoggedInUser")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayerWithId provides a mock function with given fields: ctx, id, reqEditors
func (_m *ClientInterface) GetPlayerWithId(ctx context.Context, id string, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPlayerWithId")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, id, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, id, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, id, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecentStatistics provides a mock function with given fields: ctx, reqEditors
func (_m *ClientInterface) GetRecentStatistics(ctx context.Context, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRecentStatistics")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListActiveGames provides a mock function with given fields: ctx, reqEditors
func (_m *ClientInterface) ListActiveGames(ctx context.Context, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListActiveGames")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlayers provides a mock function with given fields: ctx, params, reqEditors
func (_m *ClientInterface) ListPlayers(ctx context.Context, params *oapiprivate.ListPlayersParams, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPlayers")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oapiprivate.ListPlayersParams, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, params, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oapiprivate.ListPlayersParams, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, params, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oapiprivate.ListPlayersParams, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, params, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartGame provides a mock function with given fields: ctx, body, reqEditors
func (_m *ClientInterface) StartGame(ctx context.Context, body oapiprivate.StartGameJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartGame")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.StartGameJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapiprivate.StartGameJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapiprivate.StartGameJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartGameWithBody provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *ClientInterface) StartGameWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartGameWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlayerWithId provides a mock function with given fields: ctx, id, body, reqEditors
func (_m *ClientInterface) UpdatePlayerWithId(ctx context.Context, id string, body oapiprivate.UpdatePlayerWithIdJSONRequestBody, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePlayerWithId")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.UpdatePlayerWithIdJSONRequestBody, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, id, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, oapiprivate.UpdatePlayerWithIdJSONRequestBody, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, id, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, oapiprivate.UpdatePlayerWithIdJSONRequestBody, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, id, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlayerWithIdWithBody provides a mock function with given fields: ctx, id, contentType, body, reqEditors
func (_m *ClientInterface) UpdatePlayerWithIdWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...oapiprivate.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePlayerWithIdWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, id, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, id, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, ...oapiprivate.RequestEditorFn) error); ok {
		r1 = rf(ctx, id, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientInterface creates a new instance of ClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientInterface {
	mock := &ClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
