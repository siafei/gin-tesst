package model

type Users struct {
	*Model
	NickName string `json:"nick_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status string `gorm:"default:1" json:"status"`
}

func (m Users)TableName() string {
	return "blog_users"
}