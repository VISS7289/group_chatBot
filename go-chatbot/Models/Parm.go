package Models

//定义请求参数结构体
type ParmRegister struct {
	Username   string `json:"username"`    //用户名
	Password   string `json:"password"`    //密码
	RePassword string `json:"re_password"` //确认密码
	Email      string `json:"email"`       //邮箱
	VerifiCode string `json:"verifiCode"`  //验证码
}

type VerifiEmail struct {
	To       string `json:"to"`       //用户的邮箱地址
	Username string `json:"username"` //用户名
}

type VerifiExam struct {
	VerifiCode string `json:"verifiCode"`
	Username   string `json:"username"`
}

type ParmLogin struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
}

type ParmChat struct {
	ChatContent string `json:"chat_content"`
}

type ParmReturnChat struct {
	ChatContent string `json:"chat_content"`
}

type ParmQueryUserDetial struct {
	ID string `json:"id"`
}

type ParmToken struct {
	AToken string `json:"a_token"`
	RToken string `json:"r_token"`
}

type ParmSerchUsername struct {
	Username string `json:"username"`
	UserId   string `json:"user_id"`
}

type ParmSerchIsFriend struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

type ParmFriendRequest struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
	Msg      string `json:"msg"`
}

type ParmChangeImg struct {
	UserId string `json:"user_id"`
	Img    string `json:"img_base64"`
}

type ParmChange struct {
	UserId string `json:"user_id"`
	Data   string `json:"data"`
	Type   string `json:"type"`
	Psw    string `json:"psw"`
}

var JWTSercet = []byte("7289978VISS")
