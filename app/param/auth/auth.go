package auth

type Login struct {
	UserName string `form:"user_name" json:"user_name" xml:"user_name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}


type Register struct {
	UserName string `form:"user_name" json:"user_name" xml:"user_name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	NickName string `form:"nick_name" json:"nick_name" xml:"nick_name" binding:"required"`
}
