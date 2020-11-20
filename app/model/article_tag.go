package model

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId uint32	`json:"tag_id"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
}

func (m ArticleTag)TableName() string {
	return "blog_article_tag"
}
