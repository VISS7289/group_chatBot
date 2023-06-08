package JWT

import (
	"fmt"
	"go-chatbot/Models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenAExpireDuration = time.Minute * 10
	TokenRExpireDuration = time.Hour * 24 * 7
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

//此函数用于为指定用户ID生成访问Token和刷新Token
func GenToken(userID string) (aToken, rToken string, err error) {
	// 用户ID转int
	uid, _ := strconv.ParseInt(userID, 10, 64)
	// 生成访问Token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		uid,
		// 签名算法
		jwt.StandardClaims{
			//过期时间
			ExpiresAt: time.Now().Add(TokenAExpireDuration).Unix(),
			//签名
			Issuer:    "my-project",
		},
	}).SignedString(Models.JWTSercet)
	// 生成刷新Token
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenRExpireDuration).Unix(),
		Issuer:    "my-project",
	}).SignedString(Models.JWTSercet)
	return aToken, rToken, err
}

//此函数用于解析JWT访问Token,获取MyClaims信息
func ParseToken(tokenString string) (*MyClaims, error) {
	// 创建MyClaims对象
	var mc = new(MyClaims)
	// 解析JWT字符串,解析完毕后将声明数据填充到mc中
	token, err := jwt.ParseWithClaims(tokenString, mc, JWTKeyFunc)
	//解析错误
	if err != nil {
		return nil, err
	}
	// 验证token有效性
	if token.Valid {
		return mc, nil
	}
	//验证失败
	return nil, Models.ErrorAuthHeaderValid
}

//此函数用于解析刷新Token字符串,获取声明信息
func ParseRToken(rToken string) (err error) {
	// 创建标准声明对象
	var mc = new(jwt.StandardClaims)
	// 解析刷新Token字符串
	token, err := jwt.ParseWithClaims(rToken, mc, JWTKeyFunc)
	//解析错误
	if err != nil {
		return err
	}
	// 验证token是否有效
	if token.Valid {
		return nil
	}
	return Models.ErrorAuthHeaderValid
}

// 刷新Token的功能函数
func RefreshToken(aToken, rToken string) (newAToken, newRtoken string, err error) {
	// 解析刷新Token字符串
	if err = ParseRToken(rToken); err != nil {
		return
	}
	// 解析、验证访问Token字符串
	var mc = new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, mc, JWTKeyFunc)
	fmt.Println(err)
	v, _ := err.(*jwt.ValidationError)
	fmt.Println(v.Errors)
	// 访问Token过期,重新生成Token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(strconv.FormatInt(mc.UserID, 10))
	}
	return
}

func JWTKeyFunc(token *jwt.Token) (interface{}, error) {
	return Models.JWTSercet, nil
}
