package Mysql

import (
	"fmt"
	"go-chatbot/Models"
)

func InsertFriend(req *Models.ParmFriendRequest, state int) (err error) {
	sqlStr := `INSERT INTO friend(user_id,friend_id,state,markName) VALUES(?,?,?,?)`
	if _, err = db.Exec(sqlStr, req.UserId, req.FriendId, state, req.Msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ChangeFriend(req *Models.ParmFriendRequest, state int) (err error) {
	sqlStr := `UPDATE friend SET state = ? , markName = ? WHERE user_id = ? AND friend_id = ?`
	if _, err = db.Exec(sqlStr, state, req.Msg, req.UserId, req.FriendId); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
