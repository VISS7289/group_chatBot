package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
	"go-chatbot/Pkg/Img"
	"go-chatbot/Pkg/JWT"
	"go-chatbot/Pkg/SnowFlake"
)

// 注册
func Register(p *Models.ParmRegister) (Models.UserInfoReturn, error, Models.ResCode) {
	// 校验用户名是否已存在
	if err := Mysql.QuaryUserByUsername(p.Username); err != nil {
		var re Models.UserInfoReturn
		return re, err, Models.CodeUserExist
	}
	//查重 debug阶段去掉该代码
	//if err := Mysql.QuaryUserByEmail(p.Email); err != nil {
	//	return "", "", err, Models.CodeEmailExist
	//}
	//生成UID
	userID := SnowFlake.GenID()
	//构造用户信息
	user := &Models.UserRegist{
		userID,
		p.Username,
		p.Password,
		p.Email,
		Img.DefaultImg,
	}
	//密码加密
	user.Password = EncryptPassword.EP(user.Password)
	//保存入数据库
	if err := Mysql.InsertUser(user); err != nil {
		var re Models.UserInfoReturn
		return re, err, Models.CodeServerBusy
	}
	//生成访问 tokens
	resAToken, resRToken, err := JWT.GenToken(user.UserID)
	info := Models.UserInfoReturn{
		UserID:   user.UserID,
		Username: user.Username,
		Img:      user.Img,
		Email:    user.Email,
		AToken:   resAToken,
		RToken:   resRToken,
	}
	return info, err, Models.CodeSuccess
}

//此函数用于校验用户名是否已存在
func RepeatUsername(username string) (error, Models.ResCode) {
	if err := Mysql.QuaryUserByUsername(username); err != nil {
		return err, Models.CodeUserExist
	}
	return nil, Models.CodeSuccess
}

//此函数用于用户登陆
func Login(p *Models.ParmLogin) (info Models.UserInfoReturn, err error) {
	user := &Models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//从数据库查询用户信息
	if err = Mysql.Login(user); err != nil {
		return
	}
	//生成访问 tokens
	resAToken, resRToken, err := JWT.GenToken(user.UserID)
	info = Models.UserInfoReturn{
		UserID:   user.UserID,
		Username: user.Username,
		Img:      user.Img,
		Email:    user.Email,
		AToken:   resAToken,
		RToken:   resRToken,
	}
	return info, err
}

//此函数用于校验用户注册时输入的验证码
func VerifiExam(p *Models.ParmRegister) Models.ResCode {
	//根据邮箱从Redis中查询验证码
	val, err := Redis.RedisFind(p.Email)
	//若输入的验证码与Redis中的一致,验证成功
	if err != nil {
		return Models.CodeVerifiNotFund
	}
	if val == p.VerifiCode {
		return Models.CodeSuccess
	} else {
		return Models.CodeVerifiErr
	}
}

//此函数用于根据用户ID查询对应用户详细信息
func QueryDetialByID(p *Models.ParmQueryUserDetial) (info Models.UserInfoDetailReturn, err error) {
	user := &Models.User{
		UserID: p.ID,
	}
	if err = Mysql.QuaryUserById(user); err != nil {
		return info, err
	}
	info = Models.UserInfoDetailReturn{
		UserID:   user.UserID,
		Username: user.Username,
		Img:      user.Img,
		Email:    user.Email,
		Gender: user.Gender,
		Birthday: user.Birthday,
		Intr: user.Intr,
		CreateTime: user.CreateTime,
		Phone: user.Phone,
	}
	return info, err
}