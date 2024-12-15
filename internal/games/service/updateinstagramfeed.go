package games

import (
	"common/models"
	"common/oapipublic"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const INSTAGRAM_BASE_URL = "https://graph.instagram.com/v21.0"

type InstagramFeedData struct {
	Data []models.InstagramMedia `json:"data,omitempty"`
}

func toQueryString(params map[string]string) string {
	var sb strings.Builder

	// Iterate over the map and append key-value pairs to the query string
	for key, value := range params {
		// Use url.QueryEscape to escape special characters
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(url.QueryEscape(key) + "=" + url.QueryEscape(value))
	}

	return sb.String()
}

func getMediaInformation(medias []models.InstagramMedia) (updatedMedias []models.InstagramMedia, err error) {
	for _, media := range medias {

		params := map[string]string{
			"access_token": os.Getenv("INSTAGRAM_ACCESS_TOKEN"),
			"fields":       "id,media_type,media_url,permalink,timestamp",
		}

		getMediaURL := INSTAGRAM_BASE_URL + "/" + media.MediaID + "?" + toQueryString(params)

		resp, err := http.Get(getMediaURL)
		if err != nil {
			return []models.InstagramMedia{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []models.InstagramMedia{}, err
		}

		jsonData := strings.Replace(string(body[:]), "+0000", "+00:00", -1)

		// Unmarshal the JSON data into the struct
		err = json.Unmarshal([]byte(jsonData), &media)
		if err != nil {
			return []models.InstagramMedia{}, err
		}

		updatedMedias = append(updatedMedias, media)
	}

	return updatedMedias, nil
}

func getLatestMedias(allResults bool) ([]models.InstagramMedia, error) {

	params := map[string]string{
		"access_token": os.Getenv("INSTAGRAM_ACCESS_TOKEN"),
	}

	if !allResults {
		now := time.Now()
		params["since"] = strconv.FormatInt((now.Add(time.Minute * -70).Unix()), 10)
	}

	listMediaURL := INSTAGRAM_BASE_URL + "/" + os.Getenv("INSTAGRAM_ACCOUNT_ID") + "/media?" + toQueryString(params)

	resp, err := http.Get(listMediaURL)
	if err != nil {
		return []models.InstagramMedia{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.InstagramMedia{}, err
	}

	// Create a struct variable to store the unmarshalled data
	var media InstagramFeedData

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &media)
	if err != nil {
		return []models.InstagramMedia{}, err
	}

	return media.Data, nil
}

func (gs GameService) UpdateInstagramFeed(ctx context.Context, input oapipublic.UpdateInstagramFeedRequestObject) (oapipublic.UpdateInstagramFeedResponseObject, error) {

	feedCount, err := gs.GameStore.GetInstagramFeedCount(ctx, nil)
	if err != nil {
		return nil, err
	}

	getAllMedia := false
	if feedCount < 1 {
		getAllMedia = true
	}

	medias, err := getLatestMedias(getAllMedia)
	if err != nil {
		return nil, err
	}

	medias, err = getMediaInformation(medias)
	if err != nil {
		return nil, err
	}

	for _, media := range medias {
		err = gs.GameStore.UpdateInstagramFeed(ctx, nil, media)
		if err != nil {
			return nil, err
		}
	}

	return oapipublic.UpdateInstagramFeed200Response{}, nil
}
