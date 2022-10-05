package Mysql

import (
	"go-chatbot/Models"
)

func InsertFriend(req *Models.ParmFriendRequest, state int) (err error) {
	sqlStr := `INSERT INTO friend(user_id,friend_id,state,markName) VALUES(?,?,?,?)`
	if _, err = db.Exec(sqlStr, req.UserId, req.FriendId, state, req.Msg); err != nil {
		return err
	}
	return nil
}

func ChangeFriend(req *Models.ParmFriendRequest, state int) (err error) {
	sqlStr := `UPDATE friend SET state = ? , markName = ? WHERE user_id = ? AND friend_id = ?`
	if _, err = db.Exec(sqlStr, state, req.Msg, req.UserId, req.FriendId); err != nil {
		return err
	}
	sqlStr = `INSERT INTO friend(user_id,friend_id,state,markName) VALUES(?,?,?,?)`
	if _, err = db.Exec(sqlStr, req.FriendId, req.UserId, state, req.Msg); err != nil {
		return err
	}
	return nil
}

func GetMyFriend(userid string, state int) ([]Models.MyFriend, error)  {
	sqlStr := `SELECT friend_id, last_time, markName FROM friend WHERE user_id = ? AND state = ?`
	var userRes []Models.MyFriend
	err := db.Select(&userRes, sqlStr, userid,state)
	return userRes, err
}

func GetMyFriendReq(userid string, state int) ([]Models.MyFriendReq, error)  {
	sqlStr := `SELECT user_id, last_time, markName FROM friend WHERE friend_id = ? AND state = ?`
	var userRes []Models.MyFriendReq
	err := db.Select(&userRes, sqlStr, userid,state)
	return userRes, err
}

func DelFriendById(userId string, friendId string) error {
	sqlStr := "DELETE FROM friend WHERE user_id = ? AND friend_id = ?"
	if _,err := db.Exec(sqlStr,userId,friendId);err!=nil{
		return err
	}
	if _,err := db.Exec(sqlStr,friendId,userId);err!=nil{
		return err
	}
	return nil
}

func RejectFriendById(userId string, friendId string) error {
	sqlStr := "DELETE FROM friend WHERE user_id = ? AND friend_id = ?"
	if _,err := db.Exec(sqlStr,userId,friendId);err!=nil{
		return err
	}
	return nil
}


