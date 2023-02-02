package model

import "time"

type ResBodyArticle struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResBodyArticleDetail struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
