package Mysql

import "go-chatbot/Models"

func InsertFriend(req *Models.ParmFriendRequest, state int) (err error) {
	sqlStr := `INSERT INTO friend(user_id,friend_id,state,markName) VALUES(?,?,?,?)`
	if _, err = db.Exec(sqlStr, req.UserId, req.FriendId, state, req.Msg); err != nil {
		return err
	}
	return nil
}
