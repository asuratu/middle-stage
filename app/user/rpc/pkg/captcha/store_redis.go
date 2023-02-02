package captcha

import (
	"errors"
	"github.com/zeromicro/go-zero/core/service"
	"middle/app/user/rpc/internal/svc"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	KeyPrefix  string
	CaptConfig *CaptConfig
}

// captConfig 验证码配置

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(svcCtx *svc.ServiceContext, key string, value string) error {

	// 验证码过期时间
	ExpireTime := s.CaptConfig.ExpireTime
	// 方便本地开发调试
	if s.CaptConfig.Mode == service.DevMode {
		ExpireTime = s.CaptConfig.DebugExpireTime
	}

	// 使用 redis 存储验证码
	if err := svcCtx.Redis.Setex(s.KeyPrefix+key, value, int(ExpireTime)); err != nil {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
