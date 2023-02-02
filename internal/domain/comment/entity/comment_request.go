package entity

type ReqBodyCreateComment struct {
	ArticleSlug string
	UserID      string
	Content     string `form:"content" json:"content" xml:"content" binding:"required"`
}

type ReqBodyUpdateComment struct {
	ID          string
	ArticleSlug string
	UserID      string
	Content     string `form:"content" json:"content" xml:"content" binding:"required"`
}
