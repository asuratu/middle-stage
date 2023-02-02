// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"sync"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/pkg/mail"
	"middle/app/user/rpc/pkg/sms"
	"middle/pkg/helpers"

	"github.com/zeromicro/go-zero/core/service"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode(svcCtx *svc.ServiceContext) *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				// 增加前缀保持数据库整洁，出问题调试时也方便
				KeyPrefix: svcCtx.Config.Name + ":verifycode:",
				svcCtx:    svcCtx,
			},
		}
	})

	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例：
//
//	verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(svcContext *svc.ServiceContext, phone string) bool {

	// 生成验证码
	code := vc.generateVerifyCode(svcContext, phone)

	// 方便本地和 API 自动测试
	if !(svcContext.Config.Mode == service.ProMode) && strings.HasPrefix(phone, svcContext.Config.VerifyCode.DebugPhonePrefix) {
		return true
	}

	// 发送短信
	return sms.NewSMS().Send(svcContext, phone, sms.Message{
		Template: svcContext.Config.Sms.Aliyun.TemplateCode,
		Data:     map[string]string{"code": code},
	})
}

// SendEmail 发送邮件验证码，调用示例：
//
//	verifycode.NewVerifyCode().SendEmail(request.Email)
func (vc *VerifyCode) SendEmail(svcContext *svc.ServiceContext, email string) error {

	// 生成验证码
	code := vc.generateVerifyCode(svcContext, email)

	// 方便本地和 API 自动测试
	if !(svcContext.Config.Mode == service.ProMode) && strings.HasSuffix(email, svcContext.Config.VerifyCode.DebugEmailSuffix) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	// 发送邮件
	mail.NewMailer().Send(svcContext, mail.Email{
		From: mail.From{
			Address: svcContext.Config.Email.From.Address,
			Name:    svcContext.Config.Email.From.Name,
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(svcContext *svc.ServiceContext, key string, answer string) bool {

	logx.Errorf("验证码 : 检查验证码 %+v", map[string]string{key: answer})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !(svcContext.Config.Mode == service.ProMode) &&
		(strings.HasSuffix(key, svcContext.Config.VerifyCode.DebugEmailSuffix) ||
			strings.HasPrefix(key, svcContext.Config.VerifyCode.DebugPhonePrefix)) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(svcContext *svc.ServiceContext, key string) string {

	// 生成随机码
	code := helpers.RandomNumber(svcContext.Config.VerifyCode.CodeLength)

	// 为方便开发，本地环境使用固定验证码
	if svcContext.Config.Mode == service.DevMode {
		code = svcContext.Config.VerifyCode.DebugCode
	}

	logx.Errorf("验证码 : 生成验证码 %+v", map[string]string{key: code})

	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
