package games

import (
	"common/oapipublic"
	"context"
)

func (gs GameService) GetInstagramFeed(ctx context.Context, input oapipublic.GetInstagramFeedRequestObject) (oapipublic.GetInstagramFeedResponseObject, error) {
	feed, err := gs.GameStore.GetLatestInstagramFeed(ctx, nil)
	if err != nil {
		return nil, err
	}

	resData := []oapipublic.InstagramMedia{}
	for _, media := range feed {
		resData = append(resData, media.ModelToAPI())
	}

	return oapipublic.GetInstagramFeed200JSONResponse{
		Feed: resData,
	}, nil
}
