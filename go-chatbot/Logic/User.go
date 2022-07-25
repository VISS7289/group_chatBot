package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
	"go-chatbot/Pkg/JWT"
	"go-chatbot/Pkg/SnowFlake"
)

func Register(p *Models.ParmRegister) (string, string, error, Models.ResCode) {
	//查重
	if err := Mysql.QuaryUserByUsername(p.Username); err != nil {
		return "", "", err, Models.CodeUserExist
	}
	//查重 debug阶段去掉该代码
	//if err := Mysql.QuaryUserByEmail(p.Email); err != nil {
	//	return "", "", err, Models.CodeEmailExist
	//}
	//生成UID
	userID := SnowFlake.GenID()
	user := &Models.UserRegist{
		userID,
		p.Username,
		p.Password,
		p.Email,
	}
	//密码加密
	user.Password = EncryptPassword.EP(user.Password)
	//保存入数据库
	if err := Mysql.InsertUser(user); err != nil {
		return "", "", err, Models.CodeServerBusy
	}
	atoken, rtoken, err := JWT.GenToken(user.UserID)
	return atoken, rtoken, err, Models.CodeSuccess
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

func VerifiExam(p *Models.ParmRegister) Models.ResCode {
	val, err := Redis.RedisFind(p.Email)
	if err != nil {
		return Models.CodeVerifiNotFund
	}
	if val == p.VerifiCode {
		return Models.CodeSuccess
	} else {
		return Models.CodeVerifiErr
	}
}
