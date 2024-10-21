// Package oapiprivate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package oapiprivate

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// AddPlayerRequestSchema defines model for AddPlayerRequestSchema.
type AddPlayerRequestSchema struct {
	Name string `json:"name"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Player defines model for Player.
type Player struct {
	CreatedAt string `json:"created_at"`

	// Id The unique identifier for the player.
	Id string `json:"id"`

	// Name The name of the player.
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`

	// UserId The unique identifier for the player's user.
	UserId string `json:"user_id"`
}

// User defines model for User.
type User struct {
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Id        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

// CurrentUserResponseSchema defines model for CurrentUserResponseSchema.
type CurrentUserResponseSchema struct {
	User User `json:"user"`
}

// ErrorResponseSchema defines model for ErrorResponseSchema.
type ErrorResponseSchema = Error

// ListPlayersParams defines parameters for ListPlayers.
type ListPlayersParams struct {
	// OwnerId The ID of the owner to filter players.
	OwnerId *string `form:"owner_id,omitempty" json:"owner_id,omitempty"`

	// Page The page number for pagination.
	Page int `form:"page" json:"page"`

	// PageSize The number of players per page.
	PageSize int `form:"pageSize" json:"pageSize"`

	// SortArrangement sort by and direction.
	SortArrangement *string `form:"sortArrangement,omitempty" json:"sortArrangement,omitempty"`
}

// UpdatePlayerWithIdJSONBody defines parameters for UpdatePlayerWithId.
type UpdatePlayerWithIdJSONBody struct {
	Name string `json:"name"`
}

// AddPlayerJSONRequestBody defines body for AddPlayer for application/json ContentType.
type AddPlayerJSONRequestBody = AddPlayerRequestSchema

// UpdatePlayerWithIdJSONRequestBody defines body for UpdatePlayerWithId for application/json ContentType.
type UpdatePlayerWithIdJSONRequestBody UpdatePlayerWithIdJSONBody
