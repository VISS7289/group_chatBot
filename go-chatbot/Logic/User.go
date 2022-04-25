package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
	"go-chatbot/Pkg/JWT"
	"go-chatbot/Pkg/SnowFlake"
)

func Register(p *Models.ParmRegister) (aToken, rToken string, err error) {
	//查重
	if err = Mysql.QuaryUserByUsername(p.Username); err != nil {
		return
	}
	//生成UID
	userID := SnowFlake.GenID()
	user := &Models.User{
		userID,
		p.Username,
		p.Password,
		"def",
		3,
	}
	//密码加密
	user.Password = EncryptPassword.EP(user.Password)
	//保存入数据库
	if err = Mysql.InsertUser(user); err != nil {
		return
	}
	return JWT.GenToken(user.UserID)
}

func Login(p *Models.ParmLogin) (aToken, rToken string, err error) {
	user := &Models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = Mysql.Login(user); err != nil {
		return
	}
	return JWT.GenToken(user.UserID)
}
