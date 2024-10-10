// Package oapiprivate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package oapiprivate

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// ErrorResponseSchema defines model for ErrorResponseSchema.
type ErrorResponseSchema = Error
