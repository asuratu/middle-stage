package svc

import (
	"fmt"
	"middle/app/mqueue/cmd/job/internal/config"

	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
			// 关键队列中的任务将被处理 60% 的时间
			// 默认队列中的任务将被处理 30% 的时间
			// 低队列中的任务将被处理 10% 的时间
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			StrictPriority: true, // strict mode! if true , will not exec low queue task
		},
	)
}
