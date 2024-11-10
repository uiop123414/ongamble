package models

import "time"

type News struct {
	ID        int64     `json:"id"`
	Name      string    `json:"article_name"`
	CreatedAt time.Time `json:"publish_at"`
}
