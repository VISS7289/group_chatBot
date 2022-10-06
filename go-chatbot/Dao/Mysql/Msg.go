package Mysql

import "go-chatbot/Models"

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
	if err := db.Get(&num, sqlStr,sendid, acceptid, 0); err != nil {
		return 0,err
	}
	return num,nil
}
