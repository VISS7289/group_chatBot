package Models

type MsgRes struct {
	State string `db:"state"`
	Time  string `db:"time"`
	Msg   string `db:"message"`
	Type  int    `db:"type"`
}
