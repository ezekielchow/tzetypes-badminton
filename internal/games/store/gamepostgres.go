package games

import (
	"common/models"
	"common/utils"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GamePostgres struct {
	Queries *database.Queries
}

func (gp GamePostgres) CreateGame(ctx context.Context, tx *pgx.Tx, toCreate models.Game) (models.Game, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgClubID, err := utils.StringToPgId(toCreate.ClubID)
	if err != nil {
		return models.Game{}, err
	}

	pgCreatedAt, err := utils.TimeToPgTimestamp(toCreate.CreatedAt)
	if err != nil {
		return models.Game{}, err
	}

	dbGame, err := queries.CreateGame(ctx, database.CreateGameParams{
		ClubID:              pgClubID,
		LeftOddPlayerName:   *toCreate.LeftOddPlayerName,
		LeftEvenPlayerName:  toCreate.LeftEvenPlayerName,
		RightOddPlayerName:  *toCreate.RightOddPlayerName,
		RightEvenPlayerName: toCreate.RightEvenPlayerName,
		GameType:            toCreate.GameType,
		ServingSide:         toCreate.ServingSide,
		CreatedAt:           pgCreatedAt,
	})

	if err != nil {
		return models.Game{}, err
	}

	game := models.Game{}
	err = game.PostgresToModel(dbGame)
	if err != nil {
		return models.Game{}, err
	}

	return game, err
}

func (gp GamePostgres) CreateGameStep(ctx context.Context, tx *pgx.Tx, toCreate models.GameStep) (models.GameStep, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgGameID, err := utils.StringToPgId(toCreate.GameID)
	if err != nil {
		return models.GameStep{}, err
	}

	pgScoreAt := pgtype.Timestamp{}
	err = pgScoreAt.Scan(toCreate.ScoreAt)
	if err != nil {
		return models.GameStep{}, err
	}

	dbGameStep, err := queries.CreateGameStep(ctx, database.CreateGameStepParams{
		GameID:              pgGameID,
		TeamLeftScore:       int32(toCreate.TeamLeftScore),
		TeamRightScore:      int32(toCreate.TeamRightScore),
		ScoreAt:             pgScoreAt,
		StepNum:             int32(toCreate.StepNum),
		CurrentServer:       toCreate.CurrentServer,
		LeftEvenPlayerName:  toCreate.LeftEvenPlayerName,
		LeftOddPlayerName:   *toCreate.LeftOddPlayerName,
		RightEvenPlayerName: toCreate.RightEvenPlayerName,
		RightOddPlayerName:  *toCreate.RightOddPlayerName,
		SyncID:              toCreate.SyncId,
		IsPaused:            int32(toCreate.IsPaused),
	})
	if err != nil {
		return models.GameStep{}, err
	}

	gameStep := models.GameStep{}
	err = gameStep.PostgresToModel(dbGameStep)
	if err != nil {
		return models.GameStep{}, err
	}

	return gameStep, nil
}

func (gp GamePostgres) DeleteGameStep(ctx context.Context, tx *pgx.Tx, id string) error {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return err
	}

	err = queries.DeleteGameStep(ctx, pgID)
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) EndGame(ctx context.Context, tx *pgx.Tx, id string, isEnded bool) error {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return err
	}

	err = queries.EndGame(ctx, database.EndGameParams{
		IsEnded: isEnded,
		ID:      pgID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) GetGame(ctx context.Context, tx *pgx.Tx, id string) (models.Game, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return models.Game{}, err
	}

	res, err := queries.GetGameWithID(ctx, pgID)
	if err != nil {
		return models.Game{}, err
	}

	game := models.Game{}
	err = game.PostgresToModel(res)
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (gp GamePostgres) GetGameSteps(ctx context.Context, tx *pgx.Tx, gameID string) ([]models.GameStep, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgGameID, err := utils.StringToPgId(gameID)
	if err != nil {
		return []models.GameStep{}, err
	}

	res, err := queries.GetGameStepsWithGameID(ctx, pgGameID)
	if err != nil {
		return []models.GameStep{}, err
	}

	steps := []models.GameStep{}
	for _, dbGameStep := range res {
		step := models.GameStep{}
		err = step.PostgresToModel(dbGameStep)
		if err != nil {
			return []models.GameStep{}, err
		}

		steps = append(steps, step)
	}

	return steps, nil
}

func (gp GamePostgres) CreateStatistic(ctx context.Context, tx *pgx.Tx, gameID string, toCreate models.GameStatistic) (models.GameStatistic, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgGameID, err := utils.StringToPgId(gameID)
	if err != nil {
		return models.GameStatistic{}, err
	}

	dbRes, err := queries.CreateGameStatistic(ctx, database.CreateGameStatisticParams{
		GameID:                          pgGameID,
		TotalGameTimeSeconds:            int32(toCreate.TotalGameTimeSeconds),
		RightConsecutivePoints:          int32(toCreate.RightConsecutivePoints),
		LeftConsecutivePoints:           int32(toCreate.LeftConsecutivePoints),
		LeftLongestPointSeconds:         int32(toCreate.LeftLongestPointSeconds),
		LeftShortestPointSeconds:        int32(toCreate.LeftShortestPointSeconds),
		RightLongestPointSeconds:        int32(toCreate.RightLongestPointSeconds),
		RightShortestPointSeconds:       int32(toCreate.RightShortestPointSeconds),
		AverageTimePerPointSeconds:      int32(toCreate.AverageTimePerPointSeconds),
		LeftAverageTimePerPointSeconds:  int32(toCreate.LeftAverageTimePerPointSeconds),
		RightAverageTimePerPointSeconds: int32(toCreate.RightAverageTimePerPointSeconds),
	})
	if err != nil {
		return models.GameStatistic{}, err
	}

	statistic := models.GameStatistic{}
	err = statistic.PostgresToModel(dbRes)
	if err != nil {
		return models.GameStatistic{}, err
	}

	return statistic, nil
}

func (gp GamePostgres) GetStatisticsWithGameId(ctx context.Context, tx *pgx.Tx, gameID string) (models.GameStatistic, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgGameID, err := utils.StringToPgId(gameID)
	if err != nil {
		return models.GameStatistic{}, err
	}

	dbRes, err := queries.GetGameStatisticsWithGameID(ctx, pgGameID)
	if err != nil {
		return models.GameStatistic{}, err
	}

	statistic := models.GameStatistic{}
	err = statistic.PostgresToModel(dbRes)
	if err != nil {
		return models.GameStatistic{}, err
	}

	return statistic, nil
}

func (gp GamePostgres) CreateOrUpdateGameHistory(ctx context.Context, tx *pgx.Tx, toCreate models.GameHistory) (models.GameHistory, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(toCreate.UserID)
	if err != nil {
		return models.GameHistory{}, err
	}

	pgGameID, err := utils.StringToPgId(toCreate.GameID)
	if err != nil {
		return models.GameHistory{}, err
	}

	pgGameStartedAt, err := utils.TimeToPgTimestamp(toCreate.GameStartedAt)
	if err != nil {
		return models.GameHistory{}, err
	}

	dbRes, err := queries.CreateOrUpdateGameHistory(ctx, database.CreateOrUpdateGameHistoryParams{
		UserID:                         pgUserID,
		GameID:                         pgGameID,
		PlayerPosition:                 toCreate.PlayerPosition,
		GameStartedAt:                  pgGameStartedAt,
		GameWonBy:                      toCreate.GameWonBy,
		TotalPoints:                    int32(toCreate.TotalPoints),
		PointsWon:                      int32(toCreate.PointsWon),
		PointsLost:                     int32(toCreate.PointsLost),
		AverageTimePerPointSeconds:     int32(toCreate.AverageTimePerPointSeconds),
		AverageTimePerPointWonSeconds:  int32(toCreate.AverageTimePerPointWonSeconds),
		AverageTimePerPointLostSeconds: int32(toCreate.AverageTimePerPointLostSeconds),
		LongestRallySeconds:            int32(toCreate.LongestRallySeconds),
		LongestRallyIsWon:              int32(toCreate.LongestRallyIsWon),
		ShortestRallySeconds:           int32(toCreate.ShortestRallySeconds),
		ShortestRallyIsWon:             int32(toCreate.ShortestRallyIsWon),
		IsGameWon:                      int32(toCreate.IsGameWon),
		TotalGameTimeSeconds:           int32(toCreate.TotalGameTimeSeconds),
	})
	if err != nil {
		return models.GameHistory{}, err
	}

	gameHistory := models.GameHistory{}
	err = gameHistory.PostgresToModel(dbRes)
	if err != nil {
		return models.GameHistory{}, err
	}

	return gameHistory, nil
}

func (gp GamePostgres) GetGameHistoryGivenUserIdAndGameId(ctx context.Context, tx *pgx.Tx, userID string, gameID string) (models.GameHistory, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(userID)
	if err != nil {
		return models.GameHistory{}, err
	}

	pgGameID, err := utils.StringToPgId(gameID)
	if err != nil {
		return models.GameHistory{}, err
	}

	dbRes, err := queries.GetGameHistoryGivenUserIdAndGameId(ctx, database.GetGameHistoryGivenUserIdAndGameIdParams{
		UserID: pgUserID,
		GameID: pgGameID,
	})
	if err != nil {
		return models.GameHistory{}, err
	}

	gameHistory := models.GameHistory{}
	err = gameHistory.PostgresToModel(dbRes)
	if err != nil {
		return models.GameHistory{}, err
	}

	return gameHistory, nil
}

func (gp GamePostgres) CreateOrUpdateGameRecentStatistic(ctx context.Context, tx *pgx.Tx, toCreate models.GameRecentStatistic) (models.GameRecentStatistic, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(toCreate.UserID)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	dbRes, err := queries.CreateOrUpdateGameRecentStatistic(ctx, database.CreateOrUpdateGameRecentStatisticParams{
		UserID:                         pgUserID,
		GameCount:                      int32(toCreate.GameCount),
		Wins:                           int32(toCreate.Wins),
		Losses:                         int32(toCreate.Losses),
		TotalPoints:                    int32(toCreate.TotalPoints),
		PointsWon:                      int32(toCreate.PointsWon),
		AverageTimePerPointSeconds:     int32(toCreate.AverageTimePerPointSeconds),
		AverageTimePerPointWonSeconds:  int32(toCreate.AverageTimePerPointWonSeconds),
		AverageTimePerPointLostSeconds: int32(toCreate.AverageTimePerPointLostSeconds),
		LongestRallySeconds:            int32(toCreate.LongestRallySeconds),
		LongestRallyIsWon:              int32(toCreate.LongestRallyIsWon),
		ShortestRallySeconds:           int32(toCreate.ShortestRallySeconds),
		ShortestRallyIsWon:             int32(toCreate.ShortestRallyIsWon),
		AverageTimePerGameSeconds:      int32(toCreate.AverageTimePerGameSeconds),
		NeedsRegenerating:              int32(toCreate.NeedsRegenerating),
	})
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	grs := models.GameRecentStatistic{}
	err = grs.PostgresToModel(dbRes)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	return grs, nil
}

func (gp GamePostgres) GetGameRecentStatisticWithUserId(ctx context.Context, tx *pgx.Tx, userID string) (models.GameRecentStatistic, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(userID)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	dbRes, err := queries.GetGameRecentStatisticWithUserId(ctx, pgUserID)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	grs := models.GameRecentStatistic{}
	err = grs.PostgresToModel(dbRes)
	if err != nil {
		return models.GameRecentStatistic{}, err
	}

	return grs, nil
}

func (gp GamePostgres) GetGameRecentStatisticThatNeedsRegeneration(ctx context.Context, tx *pgx.Tx) ([]models.GameRecentStatistic, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	dbRes, err := queries.GetGameRecentStatisticThatNeedsRegeneration(ctx)
	if err != nil {
		return []models.GameRecentStatistic{}, err
	}

	grsArray := []models.GameRecentStatistic{}
	for _, res := range dbRes {
		grs := models.GameRecentStatistic{}
		err = grs.PostgresToModel(res)
		if err != nil {
			return []models.GameRecentStatistic{}, err
		}

		grsArray = append(grsArray, grs)
	}
	return grsArray, nil
}

func (gp GamePostgres) GetMostRecentGameHistories(ctx context.Context, tx *pgx.Tx, userID string) ([]models.GameHistory, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(userID)
	if err != nil {
		return []models.GameHistory{}, err
	}

	dbRes, err := queries.GetMostRecentGameHistories(ctx, pgUserID)
	if err != nil {
		return []models.GameHistory{}, err
	}

	ghArray := []models.GameHistory{}
	for _, res := range dbRes {
		gh := models.GameHistory{}
		err = gh.PostgresToModel(res)
		if err != nil {
			return []models.GameHistory{}, err
		}

		ghArray = append(ghArray, gh)
	}
	return ghArray, nil
}

func (gp GamePostgres) GetGameStepsGivenGameIds(ctx context.Context, tx *pgx.Tx, gameIDs []string) ([]models.GameStep, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	ids := []pgtype.UUID{}
	for _, gameId := range gameIDs {
		id, err := utils.StringToPgId(gameId)
		if err != nil {
			return []models.GameStep{}, nil
		}

		ids = append(ids, id)
	}

	dbRes, err := queries.GetGameStepsGivenGameIds(ctx, ids)
	if err != nil {
		return []models.GameStep{}, nil
	}

	gsArray := []models.GameStep{}
	for _, res := range dbRes {
		gs := models.GameStep{}
		err = gs.PostgresToModel(res)
		if err != nil {
			return []models.GameStep{}, err
		}

		gsArray = append(gsArray, gs)
	}
	return gsArray, nil
}

func (gp GamePostgres) GetAbandonedGames(ctx context.Context, tx *pgx.Tx) ([]string, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	dbRes, err := queries.GetAbandonedGames(ctx)
	if err != nil {
		return []string{}, nil
	}

	ids := []string{}
	for _, pgID := range dbRes {
		id, err := uuid.FromBytes(pgID.Bytes[:])
		if err != nil {
			return []string{}, nil
		}

		ids = append(ids, id.String())
	}

	return ids, nil
}

func (gp GamePostgres) EndGames(ctx context.Context, tx *pgx.Tx, ids []string) error {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgIds := []pgtype.UUID{}
	for _, id := range ids {
		pgId, err := utils.StringToPgId(id)
		if err != nil {
			return err
		}

		pgIds = append(pgIds, pgId)
	}

	err := queries.EndGames(ctx, pgIds)
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) GetActiveGames(ctx context.Context, tx *pgx.Tx, clubID string) ([]models.Game, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgClubID, err := utils.StringToPgId(clubID)
	if err != nil {
		return []models.Game{}, err
	}

	dbRes, err := queries.GetActiveGames(ctx, pgClubID)
	if err != nil {
		return []models.Game{}, err
	}

	games := []models.Game{}
	for _, row := range dbRes {
		game := models.Game{}
		err = game.PostgresToModel(row)
		if err != nil {
			return []models.Game{}, err
		}

		games = append(games, game)
	}

	return games, nil
}

func (gp GamePostgres) GetPlayedGames(ctx context.Context, tx *pgx.Tx, userID string, sortIsGameWon string, sortGameCreatedAt string, offset int, limit int) ([]database.GetPlayedGamesRow, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgUserID, err := utils.StringToPgId(userID)
	if err != nil {
		return []database.GetPlayedGamesRow{}, err
	}

	dbRes, err := queries.GetPlayedGames(ctx, database.GetPlayedGamesParams{
		UserID:            pgUserID,
		SortIsGameWon:     sortIsGameWon,
		SortGameCreatedAt: sortGameCreatedAt,
		OffsetCount:       int32(offset),
		LimitCount:        int32(limit),
	})
	if err != nil {
		return []database.GetPlayedGamesRow{}, err
	}

	return dbRes, nil
}

func (gp GamePostgres) GetClubGames(ctx context.Context, tx *pgx.Tx, clubID string, sortGameCreatedAt string, offset int, limit int) ([]database.GetClubGamesRow, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgClubID, err := utils.StringToPgId(clubID)
	if err != nil {
		return []database.GetClubGamesRow{}, err
	}

	dbRes, err := queries.GetClubGames(ctx, database.GetClubGamesParams{
		ClubID:            pgClubID,
		SortGameCreatedAt: sortGameCreatedAt,
		OffsetCount:       int32(offset),
		LimitCount:        int32(limit),
	})
	if err != nil {
		return []database.GetClubGamesRow{}, err
	}

	return dbRes, nil
}

func (gp GamePostgres) UpdateInstagramFeed(ctx context.Context, tx *pgx.Tx, media models.InstagramMedia) error {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	pgPostedAt, err := utils.TimeToPgTimestamp(media.PostedAt)
	if err != nil {
		return err
	}

	err = queries.UpdateInstagramFeed(ctx, database.UpdateInstagramFeedParams{
		MediaID:   media.MediaID,
		MediaType: media.MediaType,
		MediaUrl:  media.MediaUrl,
		Permalink: media.MediaUrl,
		PostedAt:  pgPostedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) GetInstagramFeedCount(ctx context.Context, tx *pgx.Tx) (int64, error) {
	queries := gp.Queries
	if tx != nil {
		queries = queries.WithTx(*tx)
	}

	return queries.GetInstagramFeedCount(ctx)
}
