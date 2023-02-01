package captcha

import (
	"errors"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	KeyPrefix  string
	CaptConfig *CaptConfig
}

// captConfig 验证码配置

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {

	ExpireTime := time.Minute * time.Duration(s.CaptConfig.ExpireTime)
	// 方便本地开发调试
	if {
		ExpireTime = time.Minute * time.Duration(s.CaptConfig.DebugExpireTime)
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
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
