package Email

import (
	"crypto/tls"
	"fmt"
	"go-chatbot/Settings"
	gomail "gopkg.in/mail.v2"
)

var m *gomail.Message
var d *gomail.Dialer

//此函数用于初始化邮件发送功能
func EmailInit(email Settings.EmailConfig) {
	// 初始化gomail.Message
	m = gomail.NewMessage()
	// 设置email头
	m.SetHeader("From", email.From) // 发件人
	// 初始化gomail.Dialer
	d = gomail.NewDialer(
		email.Host,
		email.Port,
		email.From,
		email.Password,
	)
	// 跳过TLS验证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

// 发送邮件的函数
func SendEmail(to string,subject string,user string,verificationCode string)(err error)  {
	// 构造邮件内容
	message := `
    <p> 您好， %s,</p>
	
		<p style="text-indent:2em">您的验证码为：%s。</p> 

		<p style="text-indent:2em">验证码有效期为五分钟。</p> 

		<p style="text-indent:2em">感谢您的使用！</p>
	`
	// 设置收件人
	m.SetHeader("To", to)
	// 设置主题
	m.SetHeader("Subject", subject)                    // 邮件主题
	//设置HTML格式正文,并传入用户名和验证码
	m.SetBody("text/html", fmt.Sprintf(message, user,verificationCode))
	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

