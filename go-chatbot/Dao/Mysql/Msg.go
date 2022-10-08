package Mysql

import (
	"database/sql"
	"go-chatbot/Models"
)

func GetNewMsgOne(sendid string, acceptid string) (Models.MsgRes, error) {
	sqlStr := `SELECT state, time, message, type FROM singleMessage WHERE send_id = ? AND accept_id = ? ORDER BY time DESC`
	var info Models.MsgRes
	if err := db.Get(&info, sqlStr, sendid, acceptid); err != nil {
		return info, err
	}
	return info, nil
}

func GetUnKnowMsgNum(sendid string, acceptid string) (num int, err error) {
	sqlStr := `SELECT COUNT(id) FROM singleMessage WHERE send_id = ? AND accept_id = ? AND state=?`
	if err := db.Get(&num, sqlStr, sendid, acceptid, 1); err != nil {
		return 0, err
	}
	return num, nil
}

func GetOldMsgF(p *Models.ParmGetOldChat) ([]Models.MsgOldF, error) {
	//sqlStr := `SELECT send_id, accept_id, time, message, type FROM singleMessage WHERE (send_id = ? AND accept_id = ?) OR (send_id = ? AND accept_id = ?) ORDER BY time DESC LIMIT ?,?`
	sqlStr := `SELECT send_id, id, time, message, type FROM singleMessage WHERE (send_id = ? AND accept_id = ?) OR (send_id = ? AND accept_id = ?) ORDER BY time LIMIT ?,?`
	var info []Models.MsgOldF
	if err := db.Select(&info, sqlStr, p.UserId, p.FriendId, p.FriendId, p.UserId, p.NowPage*p.MaxPage, (p.NowPage+1)*p.MaxPage); err != nil {
		return info, err
	}
	sqlStr = `UPDATE singleMessage SET state = ? WHERE id = ? AND send_id = ?`
	for i:=0;i<len(info);i++ {
		if _, err := db.Exec(sqlStr, 0, info[i].Id, p.FriendId); err != nil {
			if err != sql.ErrNoRows{
				return info,err
			}
		}
	}
	return info, nil
}

func InsertMsg(msg Models.NewMsg) error {
	sqlStr := `INSERT INTO singleMessage(send_id,accept_id,state,message,type) VALUES(?,?,?,?,?)`
	if _, err := db.Exec(sqlStr, msg.SendId, msg.AcceptId, msg.State, msg.Msg, msg.Type); err != nil {
		return err
	}
	return nil
}
