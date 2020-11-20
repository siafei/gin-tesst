package model

type Users struct {
	*Model
	NickName string `json:"nick_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status string `json:"status"`
	CreatedBy string `json:"created_by"`
}

func (m Users)TableName() string {
	return "blog_users"
}