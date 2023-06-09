package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
	"strconv"
)

// 更新指定用户的头像图片
func ChangeUserImg(p *Models.ParmChangeImg) error {
	return Mysql.ChangeImgByID(p.Img, p.UserId)
}

// 此函数用于更改昵称
func ChangeNick(p *Models.ParmFriendRequest) error {
	return Mysql.ChangeNick(p)
}

// 此函数用于更改用户信息
func ChangeUser(p *Models.ParmChange) error {
	if p.Type == "username" {
		//校验用户名长度
		if len(p.Data) < 4 {
			return Models.ErrorInvalidParm
		}
		//校验用户名唯一性
		if err, _ := RepeatUsername(p.Data); err != nil {
			return err
		}
		//修改用户名
		return Mysql.ChangeNameByID(p.Data, p.UserId)
	}
	if p.Type == "password" {
		//密码加密
		p.Data = EncryptPassword.EP(p.Data)
		return Mysql.ChangePswByID(p.Data, p.UserId)
	}
	//修改邮箱
	if p.Type == "email" {
		return Mysql.ChangeEmailByID(p.Data, p.UserId)
	}
	//修改用户头像
	if p.Type == "imgBase64" {
		return Mysql.ChangeImgByID(p.Data, p.UserId)
	}
	//修改性别
	if p.Type == "gender" {
		gender, _ := strconv.Atoi(p.Data)
		return Mysql.ChangeGenderByID(gender, p.UserId)
	}
	//修改生日
	if p.Type == "birthday" {
		return Mysql.ChangeBirthdayByID(p.Data, p.UserId)
	}
	//修改签名
	if p.Type == "intr" {
		return Mysql.ChangeIntrByID(p.Data, p.UserId)
	}
	//修改手机
	if p.Type == "phone" {
		return Mysql.ChangePhoneByID(p.Data, p.UserId)
	}
	return Models.ErrorInvalidParm
}
