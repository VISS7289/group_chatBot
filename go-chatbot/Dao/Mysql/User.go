package Mysql

import (
	"database/sql"
	"go-chatbot/Models"
	"go-chatbot/Pkg/EncryptPassword"
)

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

func InsertUser(user *Models.User) (err error) {
	sqlStr := `INSERT INTO user(user_id,username,password,email,gender) VALUES(?,?,?,?,?)`
	if _, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email, user.Gender); err != nil {
		return err
	}
	return nil
}

func Login(u *Models.User) error {
	oPassword := u.Password
	sqlStr := `SELECT user_id, username, password FROM user WHERE username = ?`
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
