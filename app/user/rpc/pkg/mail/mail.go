// Package mail 发送短信
package mail

import (
	"middle/app/user/rpc/internal/svc"
	"middle/pkg/cast"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

// NewMailer 单例模式获取
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(svcContext *svc.ServiceContext, email Email) bool {
	return mailer.Driver.Send(email, cast.ToStringMapString(svcContext.Config.Email.Smtp))
}
