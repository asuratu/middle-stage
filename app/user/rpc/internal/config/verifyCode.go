package config

type VerifyCodeConfig struct {
	CodeLength       int    `json:",default=5,options=5|6|7"` // 验证码长度
	ExpireTime       int64  `json:",default=900"`             // 验证码过期时间, 单位秒
	DebugExpireTime  int64  `json:",default=604800"`          // 调试模式下验证码过期时间, 单位秒
	DebugCode        string `json:",default=123456"`          // 调试模式下验证码
	DebugPhonePrefix string `json:",default=000"`             // 调试模式下手机号前缀
	DebugEmailSuffix string `json:",default=@test.com"`       // 调试模式下邮箱后缀
}
