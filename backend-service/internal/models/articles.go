package models

import "time"

type Article struct {
	Name        string    `json:"article_name"`
	Publish     bool      `json:"publish"`
	ReadingTime string    `json:"reading_time"`
	Username    string    `json:"username"`
	HtmlList    string    `json:"html_list"`
	Version     int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
