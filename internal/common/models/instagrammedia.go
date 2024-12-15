package models

import "time"

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
