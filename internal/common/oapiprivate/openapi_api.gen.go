// Package oapiprivate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package oapiprivate

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /dashboard)
	Dashboard(w http.ResponseWriter, r *http.Request)

	// (GET /logout)
	Logout(w http.ResponseWriter, r *http.Request)
	// List players
	// (GET /players)
	ListPlayers(w http.ResponseWriter, r *http.Request, params ListPlayersParams)

	// (POST /players/add)
	AddPlayer(w http.ResponseWriter, r *http.Request)
	// Get a player by ID
	// (GET /players/{id})
	GetPlayerWithId(w http.ResponseWriter, r *http.Request, id string)
	// Update a player by ID
	// (PUT /players/{id})
	UpdatePlayerWithId(w http.ResponseWriter, r *http.Request, id string)

	// (GET /users/current)
	GetLoggedInUser(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /dashboard)
func (_ Unimplemented) Dashboard(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /logout)
func (_ Unimplemented) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List players
// (GET /players)
func (_ Unimplemented) ListPlayers(w http.ResponseWriter, r *http.Request, params ListPlayersParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /players/add)
func (_ Unimplemented) AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a player by ID
// (GET /players/{id})
func (_ Unimplemented) GetPlayerWithId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a player by ID
// (PUT /players/{id})
func (_ Unimplemented) UpdatePlayerWithId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /users/current)
func (_ Unimplemented) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// Dashboard operation middleware
func (siw *ServerInterfaceWrapper) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Dashboard(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Logout operation middleware
func (siw *ServerInterfaceWrapper) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Logout(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListPlayers operation middleware
func (siw *ServerInterfaceWrapper) ListPlayers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPlayersParams

	// ------------- Optional query parameter "owner_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "owner_id", r.URL.Query(), &params.OwnerId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "owner_id", Err: err})
		return
	}

	// ------------- Required query parameter "page" -------------

	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "page"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Required query parameter "pageSize" -------------

	if paramValue := r.URL.Query().Get("pageSize"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "pageSize"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sortArrangement" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortArrangement", r.URL.Query(), &params.SortArrangement)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortArrangement", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListPlayers(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPlayer operation middleware
func (siw *ServerInterfaceWrapper) AddPlayer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPlayer(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPlayerWithId operation middleware
func (siw *ServerInterfaceWrapper) GetPlayerWithId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPlayerWithId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdatePlayerWithId operation middleware
func (siw *ServerInterfaceWrapper) UpdatePlayerWithId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdatePlayerWithId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetLoggedInUser operation middleware
func (siw *ServerInterfaceWrapper) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLoggedInUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/dashboard", wrapper.Dashboard)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/logout", wrapper.Logout)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/players", wrapper.ListPlayers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/players/add", wrapper.AddPlayer)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/players/{id}", wrapper.GetPlayerWithId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/players/{id}", wrapper.UpdatePlayerWithId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/current", wrapper.GetLoggedInUser)
	})

	return r
}

type CurrentUserResponseSchemaJSONResponse struct {
	User User `json:"user"`
}

type ErrorResponseSchemaJSONResponse Error

type DashboardRequestObject struct {
}

type DashboardResponseObject interface {
	VisitDashboardResponse(w http.ResponseWriter) error
}

type Dashboard204Response struct {
}

func (response Dashboard204Response) VisitDashboardResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DashboarddefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response DashboarddefaultJSONResponse) VisitDashboardResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type LogoutRequestObject struct {
}

type LogoutResponseObject interface {
	VisitLogoutResponse(w http.ResponseWriter) error
}

type Logout204Response struct {
}

func (response Logout204Response) VisitLogoutResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type LogoutdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response LogoutdefaultJSONResponse) VisitLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type ListPlayersRequestObject struct {
	Params ListPlayersParams
}

type ListPlayersResponseObject interface {
	VisitListPlayersResponse(w http.ResponseWriter) error
}

type ListPlayers200JSONResponse struct {
	Pagination *struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
		TotalItems  int `json:"totalItems"`
		TotalPages  int `json:"totalPages"`
	} `json:"pagination,omitempty"`
	Players *[]Player `json:"players,omitempty"`
}

func (response ListPlayers200JSONResponse) VisitListPlayersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AddPlayerRequestObject struct {
	Body *AddPlayerJSONRequestBody
}

type AddPlayerResponseObject interface {
	VisitAddPlayerResponse(w http.ResponseWriter) error
}

type AddPlayer201JSONResponse Player

func (response AddPlayer201JSONResponse) VisitAddPlayerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type AddPlayerdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response AddPlayerdefaultJSONResponse) VisitAddPlayerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetPlayerWithIdRequestObject struct {
	Id string `json:"id"`
}

type GetPlayerWithIdResponseObject interface {
	VisitGetPlayerWithIdResponse(w http.ResponseWriter) error
}

type GetPlayerWithId200JSONResponse Player

func (response GetPlayerWithId200JSONResponse) VisitGetPlayerWithIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPlayerWithIddefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetPlayerWithIddefaultJSONResponse) VisitGetPlayerWithIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type UpdatePlayerWithIdRequestObject struct {
	Id   string `json:"id"`
	Body *UpdatePlayerWithIdJSONRequestBody
}

type UpdatePlayerWithIdResponseObject interface {
	VisitUpdatePlayerWithIdResponse(w http.ResponseWriter) error
}

type UpdatePlayerWithId200JSONResponse Player

func (response UpdatePlayerWithId200JSONResponse) VisitUpdatePlayerWithIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetLoggedInUserRequestObject struct {
}

type GetLoggedInUserResponseObject interface {
	VisitGetLoggedInUserResponse(w http.ResponseWriter) error
}

type GetLoggedInUser200JSONResponse struct {
	CurrentUserResponseSchemaJSONResponse
}

func (response GetLoggedInUser200JSONResponse) VisitGetLoggedInUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetLoggedInUserdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetLoggedInUserdefaultJSONResponse) VisitGetLoggedInUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /dashboard)
	Dashboard(ctx context.Context, request DashboardRequestObject) (DashboardResponseObject, error)

	// (GET /logout)
	Logout(ctx context.Context, request LogoutRequestObject) (LogoutResponseObject, error)
	// List players
	// (GET /players)
	ListPlayers(ctx context.Context, request ListPlayersRequestObject) (ListPlayersResponseObject, error)

	// (POST /players/add)
	AddPlayer(ctx context.Context, request AddPlayerRequestObject) (AddPlayerResponseObject, error)
	// Get a player by ID
	// (GET /players/{id})
	GetPlayerWithId(ctx context.Context, request GetPlayerWithIdRequestObject) (GetPlayerWithIdResponseObject, error)
	// Update a player by ID
	// (PUT /players/{id})
	UpdatePlayerWithId(ctx context.Context, request UpdatePlayerWithIdRequestObject) (UpdatePlayerWithIdResponseObject, error)

	// (GET /users/current)
	GetLoggedInUser(ctx context.Context, request GetLoggedInUserRequestObject) (GetLoggedInUserResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// Dashboard operation middleware
func (sh *strictHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	var request DashboardRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Dashboard(ctx, request.(DashboardRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Dashboard")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DashboardResponseObject); ok {
		if err := validResponse.VisitDashboardResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Logout operation middleware
func (sh *strictHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var request LogoutRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Logout(ctx, request.(LogoutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Logout")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LogoutResponseObject); ok {
		if err := validResponse.VisitLogoutResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListPlayers operation middleware
func (sh *strictHandler) ListPlayers(w http.ResponseWriter, r *http.Request, params ListPlayersParams) {
	var request ListPlayersRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListPlayers(ctx, request.(ListPlayersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListPlayers")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListPlayersResponseObject); ok {
		if err := validResponse.VisitListPlayersResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// AddPlayer operation middleware
func (sh *strictHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	var request AddPlayerRequestObject

	var body AddPlayerJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.AddPlayer(ctx, request.(AddPlayerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AddPlayer")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(AddPlayerResponseObject); ok {
		if err := validResponse.VisitAddPlayerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPlayerWithId operation middleware
func (sh *strictHandler) GetPlayerWithId(w http.ResponseWriter, r *http.Request, id string) {
	var request GetPlayerWithIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPlayerWithId(ctx, request.(GetPlayerWithIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPlayerWithId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPlayerWithIdResponseObject); ok {
		if err := validResponse.VisitGetPlayerWithIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdatePlayerWithId operation middleware
func (sh *strictHandler) UpdatePlayerWithId(w http.ResponseWriter, r *http.Request, id string) {
	var request UpdatePlayerWithIdRequestObject

	request.Id = id

	var body UpdatePlayerWithIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdatePlayerWithId(ctx, request.(UpdatePlayerWithIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdatePlayerWithId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdatePlayerWithIdResponseObject); ok {
		if err := validResponse.VisitUpdatePlayerWithIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetLoggedInUser operation middleware
func (sh *strictHandler) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	var request GetLoggedInUserRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetLoggedInUser(ctx, request.(GetLoggedInUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetLoggedInUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetLoggedInUserResponseObject); ok {
		if err := validResponse.VisitGetLoggedInUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
