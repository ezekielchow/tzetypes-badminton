package games

import (
	"common/oapiprivate"
	"context"
)

func (gs GameService) DeleteGameSteps(ctx context.Context, input oapiprivate.DeleteGameStepsRequestObject) (oapiprivate.DeleteGameStepsResponseObject, error) {
	tx, err := gs.PgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	for _, id := range *input.Body {

		err = gs.GameStore.DeleteGameStep(ctx, &tx, id)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return oapiprivate.DeleteGameSteps200Response{}, nil
}
