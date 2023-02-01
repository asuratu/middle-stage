package config

import (
	"middle/pkg/captcha"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpcConf zrpc.RpcClientConf
	CaptConfig  captcha.CaptConfig
}
