package Mysql

import (
	"database/sql"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
)

// 查询指定用户名的用户是否存在
func QuaryUserByUsername(username string) error {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return Models.ErrorUserExit
	}
	return nil
}

// 根据用户ID查询用户信息
func QuaryUserById(u *Models.User) error {
	sqlStr := `SELECT username, email, imgBase64, gender, birthday, intr, creat_time, phone FROM user WHERE user_id = ?`
	err := db.Get(u, sqlStr, u.UserID)
	if err == sql.ErrNoRows {
		return Models.ErrorUserNotExit
	}
	if err != nil {
		return err
	}
	return nil
}

// 根据用户ID查询用户密码是否正确
func QuaryPswById(psw string, id string) error {
	var newpsw string
	sqlStr := `SELECT password FROM user WHERE user_id = ?`
	if err := db.Get(newpsw, sqlStr, id); err != nil {
		return err
	}
	if newpsw != EncryptPassword.EP(psw) {
		return Models.ErrorWrongPassword
	}
	return nil
}

// 查询指定邮箱的用户是否存在
func QuaryUserByEmail(email string) error {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE email = ?`
	var count int
	if err := db.Get(&count, sqlStr, email); err != nil {
		return err
	}
	if count > 0 {
		return Models.ErrorEmailExit
	}
	return nil
}

// 根据用户ID查询用户的用户名和头像
func GetNameAndImgById(userId string) (Models.UserShortInfo, error) {
	sqlStr := `SELECT username, imgBase64 FROM user WHERE user_id = ?`
	var info Models.UserShortInfo
	if err := db.Get(&info, sqlStr, userId); err != nil {
		return info, err
	}
	return info, nil
}

// 根据用户ID查询用户的头像信息
func GetImgById(userId string) (string, error) {
	sqlStr := `SELECT  imgBase64 FROM user WHERE user_id = ?`
	var info string
	if err := db.Get(&info, sqlStr, userId); err != nil {
		return info, err
	}
	return info, nil
}

//func InsertUser(user *Models.User) (err error) {
//	sqlStr := `INSERT INTO user(user_id,username,password,email,gender) VALUES(?,?,?,?,?)`
//	if _, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email, user.Gender); err != nil {
//		return err
//	}
//	return nil
//}

// 将用户注册信息插入到数据库中
func InsertUser(user *Models.UserRegist) (err error) {
	sqlStr := `INSERT INTO user(user_id,username,password,email,imgBase64) VALUES(?,?,?,?,?)`
	if _, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email, user.Img); err != nil {
		return err
	}
	return nil
}

// 验证用户登录信息是否正确
func Login(u *Models.User) error {
	oPassword := u.Password
	sqlStr := `SELECT user_id, username, password, email, imgBase64 FROM user WHERE username = ?`
	err := db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return Models.ErrorUserNotExit
	}
	if err != nil {
		return err
	}
	if u.Password != EncryptPassword.EP(oPassword) {
		return Models.ErrorWrongPassword
	}
	return nil
}
