// Package mail 发送短信
package mail

import (
	"middle/app/user/rpc/internal/config"
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
	Config map[string]string
}

var once sync.Once
var internalMailer *Mailer

// NewMailer 单例模式获取
func NewMailer(c config.Config) *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
			Config: cast.ToStringMapString(c.Email.Smtp),
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, mailer.Config)
}
