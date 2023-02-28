package svc

import (
	"errors"
	"middle/app/user/rpc/pkg/es"

	"middle/app/user/rpc/internal/config"
	"middle/app/user/rpc/model"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Redis  *redis.Redis

	AsynqClient *asynq.Client
	EsClient    *elastic.Client
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化缓存
	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), errors.New("cache not found"))

	// 初始化redis
	rs := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	// 初始化数据库
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(sqlConn, c.Cache),
		Cache:       ca,
		Redis:       rs,
		AsynqClient: newAsynqClient(c), // 初始化异步队列
		EsClient:    es.Boot(c).Client, // 初始化es
	}
}
