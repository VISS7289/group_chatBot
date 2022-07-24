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

