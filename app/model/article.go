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
	//Tabler 接口来更改默认表名
	return "blog_article"
}