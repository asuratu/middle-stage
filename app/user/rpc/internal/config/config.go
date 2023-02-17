package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// 框架配置
	zrpc.RpcServerConf
	Cache cache.CacheConf
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
	JwtAuth JwtConfig
	// 业务配置
	App        AppConfig
	CaptConfig CaptConfig
	Sms        SmsConfig
	Email      EmailConfig
	VerifyCode VerifyCodeConfig
}
