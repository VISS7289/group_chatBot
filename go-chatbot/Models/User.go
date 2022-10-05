package Models

type User struct {
	UserID     string `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Email      string `db:"email"`
	Img        string `db:"imgBase64"`
	Gender     int    `db:"gender"`
	Birthday   string `db:"birthday"`
	Intr       string `db:"intr"`
	CreateTime string `db:"creat_time"`
	Phone      string `db:"phone"`
	//RegistDay
	//BriefInt
	//HeadPhoto
}

type UserRegist struct {
	UserID   string  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Img      string `db:"imgBase64"`
}

type UserInfoReturn struct {
	UserID   string  `db:"user_id"`
	Username string `db:"username"`
	Img      string `db:"imgBase64"`
	Email    string `db:"email"`
	AToken   string
	RToken   string
}

type UserInfoDetailReturn struct {
	UserID     string  `db:"user_id"`
	Username   string `db:"username"`
	Email      string `db:"email"`
	Img        string `db:"imgBase64"`
	Gender     int    `db:"gender"`
	Birthday   string `db:"birthday"`
	Intr       string `db:"intr"`
	CreateTime string `db:"creat_time"`
	Phone      string `db:"phone"`
}

type UserShortInfo struct {
	Username   string `db:"username"`
	Img        string `db:"imgBase64"`
}
