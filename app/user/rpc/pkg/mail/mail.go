// Package mail 发送短信
package mail

import (
	"middle/app/user/api/internal/svc"
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

type ConfigEmail struct {
	Smtp struct {
		Host     string `json:",default=localhost"`
		Port     string `json:",default=1025"`
		Username string
		Password string
	}
	From struct {
		Address string `json:",default=gohub@example.com"`
		Name    string `json:",default=Gohub"`
	}
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
