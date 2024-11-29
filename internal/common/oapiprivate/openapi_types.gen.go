// Package oapiprivate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapiprivate

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for CreateOrUpdateGameHistoryRequestSchemaPlayerPosition.
const (
	LeftEvenPlayer  CreateOrUpdateGameHistoryRequestSchemaPlayerPosition = "left_even_player"
	LeftOddPlayer   CreateOrUpdateGameHistoryRequestSchemaPlayerPosition = "left_odd_player"
	RightEvenPlayer CreateOrUpdateGameHistoryRequestSchemaPlayerPosition = "right_even_player"
	RightOddPlayer  CreateOrUpdateGameHistoryRequestSchemaPlayerPosition = "right_odd_player"
)

// Defines values for GameStartRequestSchemaGameType.
const (
	Doubles GameStartRequestSchemaGameType = "doubles"
	Singles GameStartRequestSchemaGameType = "singles"
)

// Defines values for GameStartRequestSchemaServingSide.
const (
	LeftEven  GameStartRequestSchemaServingSide = "left_even"
	RightEven GameStartRequestSchemaServingSide = "right_even"
)

// AddGameStepsRequestSchema defines model for AddGameStepsRequestSchema.
type AddGameStepsRequestSchema struct {
	Steps []GameStep `json:"steps"`
}

// AddPlayerRequestSchema defines model for AddPlayerRequestSchema.
type AddPlayerRequestSchema struct {
	Name string `json:"name"`
}

// CreateOrUpdateGameHistoryRequestSchema defines model for CreateOrUpdateGameHistoryRequestSchema.
type CreateOrUpdateGameHistoryRequestSchema struct {
	PlayerPosition CreateOrUpdateGameHistoryRequestSchemaPlayerPosition `json:"player_position"`
}

// CreateOrUpdateGameHistoryRequestSchemaPlayerPosition defines model for CreateOrUpdateGameHistoryRequestSchema.PlayerPosition.
type CreateOrUpdateGameHistoryRequestSchemaPlayerPosition string

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Game defines model for Game.
type Game struct {
	ClubId              string `json:"club_id"`
	CreatedAt           string `json:"created_at"`
	GameType            string `json:"game_type"`
	Id                  string `json:"id"`
	IsEnded             bool   `json:"is_ended"`
	LeftEvenPlayerName  string `json:"left_even_player_name"`
	LeftOddPlayerName   string `json:"left_odd_player_name"`
	RightEvenPlayerName string `json:"right_even_player_name"`
	RightOddPlayerName  string `json:"right_odd_player_name"`
	ServingSide         string `json:"serving_side"`
	UpdatedAt           string `json:"updated_at"`
}

// GameHistory defines model for GameHistory.
type GameHistory struct {
	CreatedAt      string `json:"created_at"`
	GameId         string `json:"game_id"`
	Id             string `json:"id"`
	PlayerPosition string `json:"player_position"`
	UpdatedAt      string `json:"updated_at"`
	UserId         string `json:"user_id"`
}

// GameRecentStatistic defines model for GameRecentStatistic.
type GameRecentStatistic struct {
	AverageTimePerGameSeconds      int    `json:"average_time_per_game_seconds"`
	AverageTimePerPointLostSeconds int    `json:"average_time_per_point_lost_seconds"`
	AverageTimePerPointSeconds     int    `json:"average_time_per_point_seconds"`
	AverageTimePerPointWonSeconds  int    `json:"average_time_per_point_won_seconds"`
	CreatedAt                      string `json:"created_at"`
	GameCount                      int    `json:"game_count"`
	Id                             string `json:"id"`
	LongestRallyIsWon              int    `json:"longest_rally_is_won"`
	LongestRallySeconds            int    `json:"longest_rally_seconds"`
	Losses                         int    `json:"losses"`
	PointsWon                      int    `json:"points_won"`
	ShortestRallyIsWon             int    `json:"shortest_rally_is_won"`
	ShortestRallySeconds           int    `json:"shortest_rally_seconds"`
	TotalPoints                    int    `json:"total_points"`
	UpdatedAt                      string `json:"updated_at"`
	UserId                         string `json:"user_id"`
	Wins                           int    `json:"wins"`
}

// GameStartRequestSchema defines model for GameStartRequestSchema.
type GameStartRequestSchema struct {
	GameType            GameStartRequestSchemaGameType    `json:"game_type"`
	LeftEvenPlayerName  string                            `json:"left_even_player_name"`
	LeftOddPlayerName   *string                           `json:"left_odd_player_name,omitempty"`
	RightEvenPlayerName string                            `json:"right_even_player_name"`
	RightOddPlayerName  *string                           `json:"right_odd_player_name,omitempty"`
	ServingSide         GameStartRequestSchemaServingSide `json:"serving_side"`
}

// GameStartRequestSchemaGameType defines model for GameStartRequestSchema.GameType.
type GameStartRequestSchemaGameType string

// GameStartRequestSchemaServingSide defines model for GameStartRequestSchema.ServingSide.
type GameStartRequestSchemaServingSide string

// GameStep defines model for GameStep.
type GameStep struct {
	CreatedAt           string  `json:"created_at"`
	CurrentServer       string  `json:"current_server"`
	GameId              string  `json:"game_id"`
	Id                  string  `json:"id"`
	LeftEvenPlayerName  string  `json:"left_even_player_name"`
	LeftOddPlayerName   string  `json:"left_odd_player_name"`
	RightEvenPlayerName string  `json:"right_even_player_name"`
	RightOddPlayerName  string  `json:"right_odd_player_name"`
	ScoreAt             string  `json:"score_at"`
	StepNum             int     `json:"step_num"`
	SyncId              *string `json:"sync_id,omitempty"`
	TeamLeftScore       int     `json:"team_left_score"`
	TeamRightScore      int     `json:"team_right_score"`
	UpdatedAt           string  `json:"updated_at"`
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

// CreateOrUpdateGameHistoryResponseSchema defines model for CreateOrUpdateGameHistoryResponseSchema.
type CreateOrUpdateGameHistoryResponseSchema struct {
	GameHistory GameHistory `json:"game_history"`
}

// CurrentUserResponseSchema defines model for CurrentUserResponseSchema.
type CurrentUserResponseSchema struct {
	User User `json:"user"`
}

// ErrorResponseSchema defines model for ErrorResponseSchema.
type ErrorResponseSchema = Error

// GetGameHistoryResponseSchema defines model for GetGameHistoryResponseSchema.
type GetGameHistoryResponseSchema struct {
	GameHistory GameHistory `json:"game_history"`
}

// GetRecentStatisticsResponseSchema defines model for GetRecentStatisticsResponseSchema.
type GetRecentStatisticsResponseSchema struct {
	GameRecentStatistics GameRecentStatistic `json:"game_recent_statistics"`
}

// StartGame201ResponseSchema defines model for StartGame201ResponseSchema.
type StartGame201ResponseSchema struct {
	Game  Game       `json:"game"`
	Steps []GameStep `json:"steps"`
}

// EndGameJSONBody defines parameters for EndGame.
type EndGameJSONBody struct {
	IsEnded *bool `json:"isEnded,omitempty"`
}

// DeleteGameStepsJSONBody defines parameters for DeleteGameSteps.
type DeleteGameStepsJSONBody = []string

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

// StartGameJSONRequestBody defines body for StartGame for application/json ContentType.
type StartGameJSONRequestBody = GameStartRequestSchema

// EndGameJSONRequestBody defines body for EndGame for application/json ContentType.
type EndGameJSONRequestBody EndGameJSONBody

// CreateOrUpdateGameHistoryJSONRequestBody defines body for CreateOrUpdateGameHistory for application/json ContentType.
type CreateOrUpdateGameHistoryJSONRequestBody = CreateOrUpdateGameHistoryRequestSchema

// AddGameStepsJSONRequestBody defines body for AddGameSteps for application/json ContentType.
type AddGameStepsJSONRequestBody = AddGameStepsRequestSchema

// DeleteGameStepsJSONRequestBody defines body for DeleteGameSteps for application/json ContentType.
type DeleteGameStepsJSONRequestBody = DeleteGameStepsJSONBody

// AddPlayerJSONRequestBody defines body for AddPlayer for application/json ContentType.
type AddPlayerJSONRequestBody = AddPlayerRequestSchema

// UpdatePlayerWithIdJSONRequestBody defines body for UpdatePlayerWithId for application/json ContentType.
type UpdatePlayerWithIdJSONRequestBody UpdatePlayerWithIdJSONBody
