// Package sms 发送短信
package sms

import (
	"middle/app/user/rpc/internal/config"
	"sync"

	"middle/pkg/cast"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
	Config map[string]string
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS(c config.Config) *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
			Config: cast.ToStringMapString(c.Sms.Aliyun),
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, sms.Config)
}
