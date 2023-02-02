package mail

import (
	"fmt"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logx.Debugf("发送邮件 ：发件详情, %+v", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),

		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		logx.Errorf("发送邮件 ：发件出错, %+v", err.Error())
		return false
	}

	logx.Debugf("发送邮件 ：发件成功")
	return true
}
