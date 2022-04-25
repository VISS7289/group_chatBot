package Models

//定义请求参数结构体
type ParmRegister struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type ParmLogin struct {
	Username string
	Password string
}

type ParmChat struct {
	ChatContent string `json:"chat_content"`
}

type ParmReturnChat struct {
	ChatContent string `json:"chat_content"`
}

var JWTSercet = []byte("7289978VISS")
