// Package captcha 处理图片验证码逻辑
package captcha

import (
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
	CaptConfig    *CaptConfig
}

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
	Mode            string  // 验证码模式
	ServiceName     string  // 服务名称
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha(captConfig *CaptConfig) *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			KeyPrefix: captConfig.ServiceName + ":captcha:",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			captConfig.Height,   // 宽
			captConfig.Width,    // 高
			captConfig.Length,   // 长度
			captConfig.MaxSkew,  // 数字的最大倾斜角度
			captConfig.DotCount, // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	// 方便本地和 API 自动测试
	if !(c.CaptConfig.Mode == "pro") && id == c.CaptConfig.TestingKey {
		return true
	}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
