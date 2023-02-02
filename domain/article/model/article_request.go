package model

type ReqBodyCreateArticle struct {
	Title  string `form:"title" json:"title" xml:"title" binding:"required"`
	Body   string `form:"body" json:"body" xml:"body" binding:"required"`
	UserID string `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
}

type ReqBodyUpdateArticle struct {
	Title  string `form:"title" json:"title" xml:"title" binding:"required"`
	Body   string `form:"body" json:"body" xml:"body" binding:"required"`
	UserID string `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
}
