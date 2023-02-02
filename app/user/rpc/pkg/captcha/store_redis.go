package captcha

import (
	"errors"

	"middle/app/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	KeyPrefix string
	svcCtx    *svc.ServiceContext
}

// captConfig 验证码配置

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {

	// 验证码过期时间
	ExpireTime := s.svcCtx.Config.CaptConfig.ExpireTime
	// 方便本地开发调试
	if s.svcCtx.Config.Mode == service.DevMode {
		ExpireTime = s.svcCtx.Config.CaptConfig.DebugExpireTime
	}

	// 使用 redis 存储验证码
	if err := s.svcCtx.Redis.Setex(s.KeyPrefix+key, value, int(ExpireTime)); err != nil {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val, err := s.svcCtx.Redis.Get(key)
	if err != nil {
		return ""
	}
	if clear {
		_, err = s.svcCtx.Redis.Del(key)
		if err != nil {
			return ""
		}
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
