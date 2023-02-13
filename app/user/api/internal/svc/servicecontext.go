package svc

import (
	"middle/app/user/api/internal/config"
	"middle/app/user/rpc/usersvc"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserCheck rest.Middleware
	UserRpc   usersvc.Usersvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usersvc.NewUsersvc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
