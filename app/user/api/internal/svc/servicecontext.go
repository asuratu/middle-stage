package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"middle/app/user/api/internal/config"
	"middle/app/user/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Usercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
	}
}
