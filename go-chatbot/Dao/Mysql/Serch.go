package Mysql

import (
	"database/sql"
	"go-chatbot/Models"
)

func SerchByUsername(Username string) ([]Models.SerchUserInfo, error) {
	//sqlStr := `SELECT username, user_id, imgBase64, email FROM user WHERE username LIKE ?`
	sqlStr := `SELECT username, user_id, imgBase64, email FROM user WHERE username LIKE CONCAT('%',?,'%')`
	var userRes []Models.SerchUserInfo
	err := db.Select(&userRes, sqlStr, Username)
	return userRes, err
}

func SerchIsFriend(userId string, friendId string) (Models.SerchIsFriend, error) {
	sqlStr := `SELECT state, markName FROM friend WHERE user_id = ? AND friend_id = ?`
	var friendRes Models.SerchIsFriend
	err := db.Get(&friendRes, sqlStr, userId, friendId)
	if err == sql.ErrNoRows {
		friendRes.State = 2
		return friendRes, nil
	}
	return friendRes, err
}



