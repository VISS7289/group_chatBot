package Mysql

import "go-chatbot/Models"

// 更新指定用户ID的用户头像信息
func ChangeImgByID(img string, id string) (err error) {
	sqlStr := `UPDATE user SET imgBase64 = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, img, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户名信息
func ChangeNameByID(name string, id string) (err error) {
	sqlStr := `UPDATE user SET username = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, name, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户密码信息
func ChangePswByID(psw string, id string) (err error) {
	sqlStr := `UPDATE user SET password = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, psw, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户电子邮件信息
func ChangeEmailByID(email string, id string) (err error) {
	sqlStr := `UPDATE user SET email = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, email, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户性别信息
func ChangeGenderByID(gender int, id string) (err error) {
	sqlStr := `UPDATE user SET gender = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, gender, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户生日信息
//UPDATE `go-chatbot-sql`.`user` SET `birthday` = '2002-06-06' WHERE `user`.`id` = 5;
func ChangeBirthdayByID(birthday string, id string) (err error) {
	sqlStr := `UPDATE user SET birthday = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, birthday, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户简介信息
func ChangeIntrByID(intr string, id string) (err error) {
	sqlStr := `UPDATE user SET intr = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, intr, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID的用户电话号码信息
func ChangePhoneByID(phone string, id string) (err error) {
	sqlStr := `UPDATE user SET phone = ? WHERE user_id = ?`
	if _, err = db.Exec(sqlStr, phone, id); err != nil {
		return err
	}
	return nil
}

// 更新指定用户ID和好友ID之间的昵称信息
func ChangeNick(p *Models.ParmFriendRequest) error {
	sqlStr := `UPDATE friend SET  markName = ? WHERE user_id = ? AND friend_id = ?`
	if _, err := db.Exec(sqlStr, p.Msg, p.UserId, p.FriendId); err != nil {
		return err
	}
	return nil
}