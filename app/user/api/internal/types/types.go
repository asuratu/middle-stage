// Code generated by goctl. DO NOT EDIT.
package types

type PhoneExistReq struct {
	Phone string `json:"phone" validate:"required,len=11"`
}

type PhoneExistReply struct {
	Exist bool `json:"exist"`
}

type CaptchaReply struct {
	CaptchaId    string `json:"captchaId"`
	CaptchaImage string `json:"captchaImage"`
}

type SendSmsCodeReq struct {
	Phone         string `json:"phone,omitempty" v:"phone @required|length:1|gt:5#请输入手机号|手机号长度非法|大于1"`
	CaptchaId     string `json:"captcha_id,omitempty" validate:"required"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" validate:"required"`
}

type SendSmsCodeReply struct {
	Success bool `json:"success"`
}
