package games

import (
	"common/models"
	"common/utils"
	"context"
	database "tzetypes-badminton/database/generated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GamePostgres struct {
	Queries *database.Queries
}

func (gp GamePostgres) CreateGame(ctx context.Context, tx *pgx.Tx, toCreate models.Game) (models.Game, error) {
	pgClubID, err := utils.StringToPgId(toCreate.ClubID)
	if err != nil {
		return models.Game{}, err
	}

	dbGame, err := gp.Queries.CreateGame(ctx, database.CreateGameParams{
		ClubID:              pgClubID,
		LeftOddPlayerName:   *toCreate.LeftOddPlayerName,
		LeftEvenPlayerName:  toCreate.LeftEvenPlayerName,
		RightOddPlayerName:  *toCreate.RightOddPlayerName,
		RightEvenPlayerName: toCreate.RightEvenPlayerName,
		GameType:            toCreate.GameType,
		ServingSide:         toCreate.ServingSide,
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

	pgGameID, err := utils.StringToPgId(toCreate.GameID)
	if err != nil {
		return models.GameStep{}, err
	}

	pgScoreAt := pgtype.Timestamp{}
	err = pgScoreAt.Scan(toCreate.ScoreAt)
	if err != nil {
		return models.GameStep{}, err
	}

	dbGameStep, err := gp.Queries.CreateGameStep(ctx, database.CreateGameStepParams{
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
	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return err
	}

	err = gp.Queries.DeleteGameStep(ctx, pgID)
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) EndGame(ctx context.Context, tx *pgx.Tx, id string, isEnded bool) error {
	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return err
	}

	err = gp.Queries.EndGame(ctx, database.EndGameParams{
		IsEnded: isEnded,
		ID:      pgID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (gp GamePostgres) GetGame(ctx context.Context, tx *pgx.Tx, id string) (models.Game, error) {

	pgID, err := utils.StringToPgId(id)
	if err != nil {
		return models.Game{}, nil
	}

	res, err := gp.Queries.GetGameWithID(ctx, pgID)
	if err != nil {
		return models.Game{}, nil
	}

	game := models.Game{}
	err = game.PostgresToModel(res)
	if err != nil {
		return models.Game{}, nil
	}

	return game, nil
}

func (gp GamePostgres) GetGameSteps(ctx context.Context, tx *pgx.Tx, gameID string) ([]models.GameStep, error) {
	pgGameID, err := utils.StringToPgId(gameID)
	if err != nil {
		return []models.GameStep{}, nil
	}

	res, err := gp.Queries.GetGameStepsWithGameID(ctx, pgGameID)
	if err != nil {
		return []models.GameStep{}, nil
	}

	steps := []models.GameStep{}
	for _, dbGameStep := range res {
		step := models.GameStep{}
		err = step.PostgresToModel(dbGameStep)
		if err != nil {
			return []models.GameStep{}, nil
		}

		steps = append(steps, step)
	}

	return steps, nil
}
