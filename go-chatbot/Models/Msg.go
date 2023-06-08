package Models

// 消息结果数据
type MsgRes struct {
	State string `db:"state"`
	Time  string `db:"time"`
	Msg   string `db:"message"`
	Type  int    `db:"type"`
}

// 历史消息数据
type MsgOldF struct {
	SendId string `db:"send_id"`
	Id     string `db:"id"`
	Type   int    `db:"type"`
	Time   string `db:"time"`
	Msg    string `db:"message"`
	Img    string `db:"imgBase64"`
}

// 新消息数据
type NewMsg struct {
	SendId   string `db:"send_id"`
	AcceptId string `db:"accept_id"`
	State    bool   `db:"state"`
	Msg      string `db:"message"`
	Type     int    `db:"type"`
}

