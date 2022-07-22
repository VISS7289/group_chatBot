package Models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Gender   int
	//Birthday string `db:"birthday"`
	//Intr string `db:"intr"`
	//ImgBlob string `db:"imgBlob"`
	//Phone
	//RegistDay
	//BriefInt
	//HeadPhoto
}

type UserRegist struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}