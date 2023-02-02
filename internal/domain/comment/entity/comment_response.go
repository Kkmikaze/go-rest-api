package entity

import "time"

type ResBodyComment struct {
	ID        string    `json:"id"`
	ArticleID string    `json:"article_id"`
	WriteBy   string    `json:"write_by"`
	Content   string    `json:"content"`
	UpdateAt  time.Time `json:"update_at"`
}
