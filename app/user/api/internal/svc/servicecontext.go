package svc

import (
	"middle/app/user/api/internal/config"
	"middle/app/user/api/internal/middleware"
	"middle/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Usercheck rest.Middleware
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
