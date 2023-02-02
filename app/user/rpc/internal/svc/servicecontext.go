package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"middle/app/user/rpc/internal/config"
	"middle/app/user/rpc/model"
	"middle/pkg/database"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Cache     cache.Cache
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	err := error(nil)
	if database.DB, err = gormc.ConnectMysql(c.Mysql); err != nil {
		logx.Error(err)
		panic(err)
	}
	database.Connect(database.DB)

	// 初始化缓存
	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), errors.New("cache not found"))

	// 初始化redis
	rs := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(database.DB, c.Cache),
		Cache:     ca,
		Redis:     rs,
	}
}
