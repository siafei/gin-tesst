package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/siafei/gin-test/global"
	"strings"
	"time"
)

type Claims struct{
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


// 根据用户的用户名和密码产生token
func GenerateToken(username ,password string)(string,error){
	//设置token有效时间
	nowTime:=time.Now()
	expireTime:=nowTime.Add(global.JwtSetting.Express)
	fmt.Println(username,password)
	claims:=Claims{
		Username:username,
		Password:password,
		StandardClaims:jwt.StandardClaims{
			// 过期时间
			ExpiresAt:expireTime.Unix(),
			// 指定token发行人
			Issuer:global.JwtSetting.Issuer,
		},
	}

	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	//参数只能[]byte  不然会报错   key is of invalid type
	jwtSecret := []byte(global.JwtSetting.Secret)
	token,err:=tokenClaims.SignedString(jwtSecret)
	return token,err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string)(*Claims,error){
	jwtSecret := []byte(global.JwtSetting.Secret)
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims!=nil{
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err

}

func GetToken(c *gin.Context) (string,error)  {
	token := strings.TrimSpace(strings.TrimLeft(c.GetHeader("authorization"),"Bearer"))
	if token=="" {
		return "",&global.Error{
			StatusCode: 500,
			Msg: "请传入正确的token",
			Code: 400,
		}
	}
	return token,nil
}
