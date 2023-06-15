package Models

//定义请求参数结构体
type ParmRegister struct {
	Username   string `json:"username"`    //用户名
	Password   string `json:"password"`    //密码
	RePassword string `json:"re_password"` //确认密码
	Email      string `json:"email"`       //邮箱
	VerifiCode string `json:"verifiCode"`  //验证码
}

// 发送验证码
type VerifiEmail struct {
	To       string `json:"to"`       //用户的邮箱地址
	Username string `json:"username"` //用户名
}

// 验证码
type VerifiExam struct {
	VerifiCode string `json:"verifiCode"`
	Email      string `json:"email"`
}

// 登录参数
type ParmLogin struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
}

// 聊天
type ParmChat struct {
	ChatContent string `json:"chat_content"`
}

// 聊天服务器返回
type ParmChatServer struct {
	Response string `json:"response"`
	History [][]string `json:"history"`
	Status int `json:"status"`
	Time string `json:"time"`
}

// 聊天
type ParmReturnChat struct {
	ChatContent string `json:"chat_content"`
}

// 查询的用户id
type ParmQueryUserDetial struct {
	ID string `json:"id"` //用户ID
}

// token
type ParmToken struct {
	AToken string `json:"a_token"` //access token
	RToken string `json:"r_token"` //refersh token
}

// 查询用户名
type ParmSerchUsername struct {
	Username string `json:"username"` //用户名
	UserId   string `json:"user_id"` //用户id
}

// 查询好友关系
type ParmSerchIsFriend struct {
	UserId   string `json:"user_id"` //用户id
	FriendId string `json:"friend_id"` //目标id
}

// 查询好友请求
type ParmFriendRequest struct {
	UserId   string `json:"user_id"` //用户id
	FriendId string `json:"friend_id"` //好友id
	Msg      string `json:"msg"` //验证消息
}

// 更改邮箱
type ParmChangeEmail struct {
	UserId     string `json:"user_id"`
	NewEmail   string `json:"n_email"`
	VerifiCode string `json:"verifiCode"`
}

// 更改密码
type ParmChangePsw struct {
	UserId     string `json:"user_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	VerifiCode string `json:"verifiCode"`
}

// 更改图片
type ParmChangeImg struct {
	UserId string `json:"user_id"`
	Img    string `json:"img_base64"`
}

// 更改数据
type ParmChange struct {
	UserId string `json:"user_id"`
	Data   string `json:"data"`
	Type   string `json:"type"`
	Psw    string `json:"psw"`
}

// 获取好友
type ParmGetFriend struct {
	UserId string `json:"user_id"`
	State  int    `json:"state"`
}

// 获取最新消息
type ParmGetNewMsgOne struct {
	SendId   string `json:"send_id"`
	AcceptId string `json:"accept_id"`
}

// 获取历史消息
type ParmGetOldChat struct {
	UserId string `json:"user_id"`
	FriendId string `json:"friend_id"`
	NowPage int `json:"now_page"`
	MaxPage int `json:"max_page"`
}

var JWTSercet = []byte("7289978VISS")
