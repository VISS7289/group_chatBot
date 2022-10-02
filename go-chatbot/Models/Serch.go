package Models

type SerchUserInfo struct {
	Username string `db:"username"`
	Img      string `db:"imgBase64"`
	Email    string `db:"email"`
}

type SerchIsFriend struct {
	State    int64  `db:"state"`
	Markname string `db:"markName"`
}
