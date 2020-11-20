package model

type Tag struct {
	*Model
	Name string `json:"name"`
	Status string `json:"status"`
}

func (m Tag)TableName() string {
	return "blog_tag"
}