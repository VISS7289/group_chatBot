package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
	"strconv"
)

func ChangeUserImg(p *Models.ParmChangeImg) error {
	return Mysql.ChangeImgByID(p.Img, p.UserId)
}

func ChangeNick(p *Models.ParmFriendRequest) error {
	return Mysql.ChangeNick(p)
}

func ChangeUser(p *Models.ParmChange) error {
	if p.Type == "username" {
		if len(p.Data) < 4 {
			return Models.ErrorInvalidParm
		}
		if err, _ := RepeatUsername(p.Data); err != nil {
			return err
		}
		return Mysql.ChangeNameByID(p.Data, p.UserId)
	}
	if p.Type == "password" {
		//查看密码是否正确
		if err := Mysql.QuaryPswById(p.Data, p.UserId); err != nil {
			return err
		}
		//密码加密
		p.Data = EncryptPassword.EP(p.Data)
		return Mysql.ChangePswByID(p.Data, p.UserId)
	}
	if p.Type == "email" {
		//查看密码是否正确
		if err := Mysql.QuaryPswById(p.Psw, p.UserId); err != nil {
			return err
		}
		return Mysql.ChangeEmailByID(p.Data, p.UserId)
	}
	if p.Type == "imgBase64" {
		return Mysql.ChangeImgByID(p.Data, p.UserId)
	}
	if p.Type == "gender" {
		gender, _ := strconv.Atoi(p.Data)
		return Mysql.ChangeGenderByID(gender, p.UserId)
	}
	if p.Type == "birthday" {
		return Mysql.ChangeBirthdayByID(p.Data, p.UserId)
	}
	if p.Type == "intr" {
		return Mysql.ChangeIntrByID(p.Data, p.UserId)
	}
	if p.Type == "phone" {
		return Mysql.ChangePhoneByID(p.Data, p.UserId)
	}
	return Models.ErrorInvalidParm
}
