package svc

import (
	"middle/app/mqueue/cmd/job/internal/config"

	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
	}
}
