package model

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	ImgUrl string `json:"img_url"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
}

func (m Article)TableName() string {
	return "blog_article"
}