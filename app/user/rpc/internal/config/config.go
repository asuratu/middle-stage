package config

import (
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
	App   App
}
