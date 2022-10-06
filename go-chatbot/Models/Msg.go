package Models

type MsgRes struct {
	State string `db:"state"`
	Time  string `db:"time"`
	Msg   string `db:"message"`
	Type  int    `db:"type"`
}

type MsgOldF struct {
	SendId string `db:"send_id"`
	Id     string `db:"id"`
	Type   int    `db:"type"`
	Time   string `db:"time"`
	Msg    string `db:"message"`
	Img    string `db:"imgBase64"`
}
