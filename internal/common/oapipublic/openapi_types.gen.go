// Package oapipublic provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package oapipublic

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// LoginRequestSchema defines model for LoginRequestSchema.
type LoginRequestSchema struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// LoginResponseSchema defines model for LoginResponseSchema.
type LoginResponseSchema struct {
	RefreshToken string `json:"refresh_token"`
	SessionToken string `json:"session_token"`
}

// RefreshTokenResponseSchema defines model for RefreshTokenResponseSchema.
type RefreshTokenResponseSchema struct {
	SessionToken string `json:"session_token"`
}

// SignupRequestSchema defines model for SignupRequestSchema.
type SignupRequestSchema struct {
	Email          openapi_types.Email `json:"email"`
	Password       string              `json:"password"`
	PasswordRepeat string              `json:"password_repeat"`
}

// ErrorResponseSchema defines model for ErrorResponseSchema.
type ErrorResponseSchema = Error

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequestSchema

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody = SignupRequestSchema
