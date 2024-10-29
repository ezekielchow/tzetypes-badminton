// Package oapipublic provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package oapipublic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetGame request
	GetGame(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// LoginWithBody request with any body
	LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Login(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RefreshToken request
	RefreshToken(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SignupWithBody request with any body
	SignupWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Signup(ctx context.Context, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetGame(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetGameRequest(c.Server, gameId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Login(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RefreshToken(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRefreshTokenRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SignupWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSignupRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Signup(ctx context.Context, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSignupRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetGameRequest generates requests for GetGame
func NewGetGameRequest(server string, gameId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "game_id", runtime.ParamLocationPath, gameId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/game/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLoginRequest calls the generic Login builder with application/json body
func NewLoginRequest(server string, body LoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewLoginRequestWithBody(server, "application/json", bodyReader)
}

// NewLoginRequestWithBody generates requests for Login with any type of body
func NewLoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRefreshTokenRequest generates requests for RefreshToken
func NewRefreshTokenRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/refresh-token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewSignupRequest calls the generic Signup builder with application/json body
func NewSignupRequest(server string, body SignupJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSignupRequestWithBody(server, "application/json", bodyReader)
}

// NewSignupRequestWithBody generates requests for Signup with any type of body
func NewSignupRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/signup")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetGameWithResponse request
	GetGameWithResponse(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*GetGameResponse, error)

	// LoginWithBodyWithResponse request with any body
	LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginResponse, error)

	LoginWithResponse(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginResponse, error)

	// RefreshTokenWithResponse request
	RefreshTokenWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RefreshTokenResponse, error)

	// SignupWithBodyWithResponse request with any body
	SignupWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SignupResponse, error)

	SignupWithResponse(ctx context.Context, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*SignupResponse, error)
}

type GetGameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Game       Game            `json:"game"`
		Statistics *GameStatistics `json:"statistics,omitempty"`
		Steps      []GameStep      `json:"steps"`
	}
	JSONDefault *ErrorResponseSchema
}

// Status returns HTTPResponse.Status
func (r GetGameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetGameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LoginResponseSchema
	JSONDefault  *ErrorResponseSchema
}

// Status returns HTTPResponse.Status
func (r LoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RefreshTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RefreshTokenResponseSchema
	JSONDefault  *ErrorResponseSchema
}

// Status returns HTTPResponse.Status
func (r RefreshTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RefreshTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SignupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *ErrorResponseSchema
}

// Status returns HTTPResponse.Status
func (r SignupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SignupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetGameWithResponse request returning *GetGameResponse
func (c *ClientWithResponses) GetGameWithResponse(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*GetGameResponse, error) {
	rsp, err := c.GetGame(ctx, gameId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetGameResponse(rsp)
}

// LoginWithBodyWithResponse request with arbitrary body returning *LoginResponse
func (c *ClientWithResponses) LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginResponse, error) {
	rsp, err := c.LoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginResponse(rsp)
}

func (c *ClientWithResponses) LoginWithResponse(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginResponse, error) {
	rsp, err := c.Login(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginResponse(rsp)
}

// RefreshTokenWithResponse request returning *RefreshTokenResponse
func (c *ClientWithResponses) RefreshTokenWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RefreshTokenResponse, error) {
	rsp, err := c.RefreshToken(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRefreshTokenResponse(rsp)
}

// SignupWithBodyWithResponse request with arbitrary body returning *SignupResponse
func (c *ClientWithResponses) SignupWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SignupResponse, error) {
	rsp, err := c.SignupWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSignupResponse(rsp)
}

func (c *ClientWithResponses) SignupWithResponse(ctx context.Context, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*SignupResponse, error) {
	rsp, err := c.Signup(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSignupResponse(rsp)
}

// ParseGetGameResponse parses an HTTP response from a GetGameWithResponse call
func ParseGetGameResponse(rsp *http.Response) (*GetGameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetGameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Game       Game            `json:"game"`
			Statistics *GameStatistics `json:"statistics,omitempty"`
			Steps      []GameStep      `json:"steps"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseLoginResponse parses an HTTP response from a LoginWithResponse call
func ParseLoginResponse(rsp *http.Response) (*LoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LoginResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseRefreshTokenResponse parses an HTTP response from a RefreshTokenWithResponse call
func ParseRefreshTokenResponse(rsp *http.Response) (*RefreshTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RefreshTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RefreshTokenResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseSignupResponse parses an HTTP response from a SignupWithResponse call
func ParseSignupResponse(rsp *http.Response) (*SignupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SignupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
