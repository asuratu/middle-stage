// Code generated by goctl. DO NOT EDIT.
package types

type PhoneExistReq struct {
	Phone string `json:"phone,omitempty" v:"phone @required|size:11#请输入手机号|手机号称长度非法"`
}

type PhoneExistReply struct {
	Exist bool `json:"exist"`
}

type SignupByPhoneReq struct {
	Name         string `json:"name"`
	Phone        string `json:"phone,omitempty" v:"phone @required|size:11#请输入手机号|手机号称长度非法"`
	Password     string `json:"password"`
	City         string `json:"city"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	SmsCode      string `json:"sms_code"`
}

type TokenReply struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type LoginByPhoneReq struct {
	Phone         string `json:"phone,omitempty" v:"phone @required|size:11#请输入手机号|手机号称长度非法"`
	Password      string `json:"password"`
	CaptchaId     string `json:"captcha_id,omitempty"`
	CaptchaAnswer string `json:"captcha_answer,omitempty"`
}

type CaptchaReply struct {
	CaptchaId    string `json:"captchaId"`
	CaptchaImage string `json:"captchaImage"`
}

type SendSmsCodeReq struct {
	Phone         string `json:"phone,omitempty" v:"phone @required|size:11#请输入手机号|手机号称长度非法"`
	CaptchaId     string `json:"captcha_id,omitempty"`
	CaptchaAnswer string `json:"captcha_answer,omitempty"`
}

type SendSmsCodeReply struct {
	Success bool `json:"success"`
}

type UserInfoReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	City         string `json:"city"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
}

type SimpleUserInfoReply struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Avatar string `json:"avatar"`
}

type UserInfoByIdReq struct {
	Id int64 `path:"id"`
}

type UserListReq struct {
	LastId   int64  `form:"last_id"`
	PageSize int64  `form:"page_size"`
	Sort     string `form:"sort"`
	Order    string `form:"order" v:"order @in:desc,asc#排序只允许填写desc或asc"`
}

type UserListReply struct {
	Users []*SimpleUserInfoReply `json:"users"`
}
