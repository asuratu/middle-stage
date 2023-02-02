// Package sms 发送短信
package sms

import (
	"sync"

	"middle/app/user/api/internal/svc"
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
}

// ConfigSms 配置文件中的短信配置
type ConfigSms struct {
	Aliyun struct {
		TemplateCode string
	}
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (sms *SMS) Send(svcContext *svc.ServiceContext, phone string, message Message) bool {
	return sms.Driver.Send(phone, message, cast.ToStringMapString(svcContext.Config.Sms.Aliyun))
}
