package models

import (
	"common/oapipublic"
	"time"
	database "tzetypes-badminton/database/generated"

	"github.com/google/uuid"
)

type InstagramMedia struct {
	ID        string
	MediaID   string    `json:"id,omitempty"`
	MediaType string    `json:"media_type,omitempty"`
	MediaUrl  string    `json:"media_url,omitempty"`
	Permalink string    `json:"permalink,omitempty"`
	PostedAt  time.Time `json:"timestamp,omitempty"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (im *InstagramMedia) PostgresToModel(fromDb database.InstagramFeed) error {
	id, err := uuid.FromBytes(fromDb.ID.Bytes[:])
	if err != nil {
		return err
	}

	im.ID = id.String()
	im.MediaID = fromDb.MediaID
	im.MediaType = fromDb.MediaType
	im.MediaUrl = fromDb.MediaUrl
	im.Permalink = fromDb.Permalink
	im.CreatedAt = fromDb.CreatedAt.Time
	im.UpdatedAt = &fromDb.UpdatedAt.Time
	return nil
}

func (im *InstagramMedia) ModelToAPI() oapipublic.InstagramMedia {
	return oapipublic.InstagramMedia{
		CreatedAt: im.CreatedAt.String(),
		Id:        im.ID,
		MediaId:   im.MediaID,
		MediaType: im.MediaType,
		MediaUrl:  im.MediaUrl,
		Permalink: im.Permalink,
		PostedAt:  im.PostedAt.String(),
		UpdatedAt: im.UpdatedAt.String(),
	}
}
