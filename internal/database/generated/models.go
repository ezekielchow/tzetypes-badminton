// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Club struct {
	ID        pgtype.UUID
	OwnerID   pgtype.UUID
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Game struct {
	ID                  pgtype.UUID
	ClubID              pgtype.UUID
	LeftOddPlayerName   *string
	LeftEvenPlayerName  string
	RightOddPlayerName  *string
	RightEvenPlayerName string
	GameType            string
	ServingSide         string
	IsEnded             bool
	CreatedAt           pgtype.Timestamp
	UpdatedAt           pgtype.Timestamp
}

type GameHistory struct {
	ID                             pgtype.UUID
	UserID                         pgtype.UUID
	GameID                         pgtype.UUID
	PlayerPosition                 string
	IsGameWon                      int32
	GameStartedAt                  pgtype.Timestamp
	GameWonBy                      string
	TotalPoints                    int32
	PointsWon                      int32
	PointsLost                     int32
	AverageTimePerPointSeconds     int32
	AverageTimePerPointWonSeconds  int32
	AverageTimePerPointLostSeconds int32
	LongestRallySeconds            int32
	LongestRallyIsWon              int32
	ShortestRallySeconds           int32
	ShortestRallyIsWon             int32
	TotalGameTimeSeconds           int32
	CreatedAt                      pgtype.Timestamp
	UpdatedAt                      pgtype.Timestamp
}

type GameRecentStatistic struct {
	ID                             pgtype.UUID
	UserID                         pgtype.UUID
	GameCount                      *int32
	Wins                           *int32
	Losses                         *int32
	TotalPoints                    *int32
	PointsWon                      *int32
	AverageTimePerPointSeconds     *int32
	AverageTimePerPointWonSeconds  *int32
	AverageTimePerPointLostSeconds *int32
	LongestRallySeconds            *int32
	LongestRallyIsWon              *int32
	ShortestRallySeconds           *int32
	ShortestRallyIsWon             *int32
	AverageTimePerGameSeconds      *int32
	NeedsRegenerating              *int32
	CreatedAt                      pgtype.Timestamp
	UpdatedAt                      pgtype.Timestamp
}

type GameStatistic struct {
	ID                              pgtype.UUID
	GameID                          pgtype.UUID
	TotalGameTimeSeconds            *int32
	RightConsecutivePoints          *int32
	LeftConsecutivePoints           *int32
	LeftLongestPointSeconds         *int32
	LeftShortestPointSeconds        *int32
	RightLongestPointSeconds        *int32
	RightShortestPointSeconds       *int32
	AverageTimePerPointSeconds      *int32
	LeftAverageTimePerPointSeconds  *int32
	RightAverageTimePerPointSeconds *int32
	CreatedAt                       pgtype.Timestamp
	UpdatedAt                       pgtype.Timestamp
}

type GameStep struct {
	ID                  pgtype.UUID
	GameID              pgtype.UUID
	TeamLeftScore       int32
	TeamRightScore      int32
	ScoreAt             pgtype.Timestamp
	StepNum             int32
	CurrentServer       string
	LeftOddPlayerName   *string
	LeftEvenPlayerName  string
	RightOddPlayerName  *string
	RightEvenPlayerName string
	IsPaused            int32
	SyncID              string
	CreatedAt           pgtype.Timestamp
	UpdatedAt           pgtype.Timestamp
}

type InstagramFeed struct {
	ID        pgtype.UUID
	MediaID   string
	MediaType string
	MediaUrl  string
	Permalink string
	PostedAt  pgtype.Timestamp
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Player struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type PlayerClub struct {
	ID        pgtype.UUID
	PlayerID  pgtype.UUID
	ClubID    pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Session struct {
	ID                    pgtype.UUID
	UserID                pgtype.UUID
	SessionToken          pgtype.UUID
	RefreshToken          pgtype.UUID
	SessionTokenExpiresAt pgtype.Timestamp
	RefreshTokenExpiresAt pgtype.Timestamp
	CreatedAt             pgtype.Timestamp
	UpdatedAt             pgtype.Timestamp
}

type User struct {
	ID          pgtype.UUID
	FirebaseUid string
	Email       string
	AccountTier string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}
