package verifycode

import (
	"errors"

	"middle/app/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	KeyPrefix string
	Config    config.Config
	Redis     *redis.Redis
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {

	ExpireTime := s.Config.VerifyCode.ExpireTime
	// 本地环境方便调试
	if s.Config.Mode == service.DevMode {
		ExpireTime = s.Config.VerifyCode.DebugExpireTime
	}

	// 使用 redis 存储验证码
	if err := s.Redis.Setex(s.KeyPrefix+key, value, int(ExpireTime)); err != nil {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val, err := s.Redis.Get(key)
	if err != nil {
		return ""
	}
	if clear {
		_, err = s.Redis.Del(key)
		if err != nil {
			return ""
		}
	}
	return val
}

// Verify 实现 verifycode.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
