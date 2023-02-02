package config

// CaptConfig 验证码配置
type CaptConfig struct {
	ExpireTime      int64   // 验证码过期时间
	DebugExpireTime int64   // 本地调试验证码过期时间
	Height          int     // 宽
	Width           int     // 高
	Length          int     // 长度
	MaxSkew         float64 // 数字的最大倾斜角度
	DotCount        int     // 干扰点数量
	TestingKey      string  // 测试用key
}
