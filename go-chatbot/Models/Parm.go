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

type ParmToken struct {
	AToken string `json:"a_token"`
	RToken string `json:"r_token"`
}

var JWTSercet = []byte("7289978VISS")
