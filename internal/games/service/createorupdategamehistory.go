package games

import (
	"common/models"
	"common/oapiprivate"
	"context"
	"database/sql"
	"log"
	"strings"
	"time"
)

func getPersonalizedStatistics(playerPosition string, game models.Game, gameSteps []models.GameStep) models.GameHistory {
	gh := models.GameHistory{}

	if len(gameSteps) == 0 {
		// Return empty GameHistory if no gameSteps provided
		return gh
	}

	first := gameSteps[0]
	last := gameSteps[len(gameSteps)-1]

	gh.GameID = game.ID
	gh.PlayerPosition = playerPosition
	gh.GameStartedAt = game.CreatedAt
	gh.TotalPoints = last.TeamLeftScore + last.TeamRightScore
	gh.TotalGameTimeSeconds = int(last.ScoreAt.Sub(first.ScoreAt).Seconds())

	if gh.TotalPoints > 0 {
		gh.AverageTimePerPointSeconds = int(last.ScoreAt.Sub(first.ScoreAt).Seconds()) / gh.TotalPoints
	} else {
		gh.AverageTimePerPointSeconds = 0
	}

	// Determine winning team
	wonBy := models.TeamSideRight
	if last.TeamLeftScore >= last.TeamRightScore {
		wonBy = models.TeamSideLeft
	}
	gh.GameWonBy = string(wonBy)

	// Did win game
	if (strings.Contains(playerPosition, "right") && strings.Contains(gh.GameWonBy, "right")) || (strings.Contains(playerPosition, "left") && strings.Contains(gh.GameWonBy, "left")) {
		gh.IsGameWon = 1
	} else {
		gh.IsGameWon = 0
	}

	// Initialize variables
	totalTimeWinningPoints := 0
	totalWinningPoints := 0
	totalTimeLosingPoints := 0
	totalLosingPoints := 0
	gh.PointsWon = 0
	gh.PointsLost = 0
	gh.LongestRallySeconds = 0
	gh.ShortestRallySeconds = int(^uint(0) >> 1) // Max int value

	for i := 1; i < len(gameSteps); i++ {
		previous := gameSteps[i-1]
		step := gameSteps[i]
		diffSeconds := int(step.ScoreAt.Sub(previous.ScoreAt).Seconds())

		isPlayerWinningPoint := (strings.Contains(playerPosition, "left") && step.TeamLeftScore > previous.TeamLeftScore) ||
			(strings.Contains(playerPosition, "right") && step.TeamRightScore > previous.TeamRightScore)

		// Track points won and lost
		if isPlayerWinningPoint {
			totalTimeWinningPoints += diffSeconds
			totalWinningPoints++
			gh.PointsWon++
		} else {
			totalTimeLosingPoints += diffSeconds
			totalLosingPoints++
			gh.PointsLost++
		}

		// Track longest rally
		if diffSeconds > gh.LongestRallySeconds {
			log.Println("helo?", isPlayerWinningPoint, playerPosition)
			gh.LongestRallySeconds = diffSeconds
			gh.LongestRallyIsWon = boolToInt(isPlayerWinningPoint)
		}

		// Track shortest rally
		if diffSeconds < gh.ShortestRallySeconds {
			gh.ShortestRallySeconds = diffSeconds
			gh.ShortestRallyIsWon = boolToInt(isPlayerWinningPoint)
		}
	}

	// Calculate average time per point
	if totalWinningPoints > 0 {
		gh.AverageTimePerPointWonSeconds = totalTimeWinningPoints / totalWinningPoints
	}
	if totalLosingPoints > 0 {
		gh.AverageTimePerPointLostSeconds = totalTimeLosingPoints / totalLosingPoints
	}

	return gh
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (gs GameService) CreateOrUpdateGameHistory(ctx context.Context, input oapiprivate.CreateOrUpdateGameHistoryRequestObject, user models.User) (oapiprivate.CreateOrUpdateGameHistoryResponseObject, error) {
	tx, err := gs.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	game, err := gs.GameStore.GetGame(ctx, &tx, input.GameId)
	if err != nil {
		return nil, err
	}

	gameSteps, err := gs.GameStore.GetGameSteps(ctx, &tx, game.ID)
	if err != nil {
		return nil, err
	}

	gh := getPersonalizedStatistics(string(input.Body.PlayerPosition), game, gameSteps)
	gh.UserID = user.ID

	gameHistory, err := gs.GameStore.CreateOrUpdateGameHistory(ctx, &tx, gh)
	if err != nil {
		return nil, err
	}

	dbGrs, err := gs.GameStore.GetGameRecentStatisticWithUserId(ctx, &tx, user.ID)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return nil, err
	}

	grs := models.GameRecentStatistic{}
	if dbGrs.ID != "" {
		grs = dbGrs
	}
	now := time.Now()
	grs.NeedsRegenerating = 1
	grs.UpdatedAt = &now
	grs.UserID = user.ID

	_, err = gs.GameStore.CreateOrUpdateGameRecentStatistic(ctx, &tx, grs)
	if err != nil {
		return nil, err
	}

	apiRes := gameHistory.ModelToAPI()

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.CreateOrUpdateGameHistory200JSONResponse{
		CreateOrUpdateGameHistoryResponseSchemaJSONResponse: oapiprivate.CreateOrUpdateGameHistoryResponseSchemaJSONResponse{
			GameHistory: apiRes,
		},
	}, nil
}
