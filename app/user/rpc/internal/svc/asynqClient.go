package svc

import (
	"middle/app/user/rpc/internal/config"

	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
