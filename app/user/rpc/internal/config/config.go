package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"middle/app/user/rpc/pkg/captcha"
	"middle/app/user/rpc/pkg/mail"
	"middle/app/user/rpc/pkg/sms"
	"middle/app/user/rpc/pkg/verifycode"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type App struct {
	Url string
}

type Config struct {
	zrpc.RpcServerConf
	Mysql gormc.Mysql
	Cache cache.CacheConf
	Redis redis.RedisConf
	App   App

	CaptConfig captcha.CaptConfig
	Sms        sms.ConfigSms
	Email      mail.ConfigEmail
	VerifyCode verifycode.ConfigVerifyCode
}
