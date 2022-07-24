package Email

import (
	"crypto/tls"
	"fmt"
	"go-chatbot/Settings"
	gomail "gopkg.in/mail.v2"
)

var m *gomail.Message
var d *gomail.Dialer

func EmailInit(email Settings.EmailConfig) {
	m = gomail.NewMessage()
	m.SetHeader("From", email.From) // 发件人

	d = gomail.NewDialer(
		email.Host,
		email.Port,
		email.From,
		email.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func SendEmail(to string,subject string,user string,verificationCode string)(err error)  {
	message := `
    <p> 您好， %s,</p>
	
		<p style="text-indent:2em">您的验证码为：%s。</p> 

		<p style="text-indent:2em">验证码有效期为五分钟。</p> 

		<p style="text-indent:2em">感谢您的使用！</p>
	`
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)                    // 邮件主题
	m.SetBody("text/html", fmt.Sprintf(message, user,verificationCode))
	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func main() {
	message := `
    <p> 您好， %s,</p>
	
		<p style="text-indent:2em">您的验证码为：%s.</p> 

		<p style="text-indent:2em">感谢您的使用！</p>
	`

	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	host := "smtp.qcloudmail.com"
	port := 465
	userName := "registe@qkgzs.com"
	password := "7289978VISSdx"

	m := gomail.NewMessage()
	m.SetHeader("From", userName) // 发件人
	// m.SetHeader("From", "alias"+"<"+userName+">") // 增加发件人别名

	m.SetHeader("To", "1526716943@qq.com") // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	//m.SetHeader("Cc", "******@qq.com")                  // 抄送，可以多个
	//m.SetHeader("Bcc", "******@qq.com")                 // 暗送，可以多个
	m.SetHeader("Subject", "Hello!")                    // 邮件主题

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", fmt.Sprintf(message, "testUser","absewasd"))


	d = gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}


