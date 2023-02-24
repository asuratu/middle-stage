package svc

import (
	"fmt"
	"middle/app/mqueue/cmd/scheduler/internal/config"
	"time"

	"github.com/hibiken/asynq"
)

// create scheduler
func newScheduler(c config.Config) *asynq.Scheduler {

	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
				if err != nil {
					fmt.Printf("Scheduler PostEnqueueFunc <<<<<<<===>>>>> err : %+v , info : %+v", err, info)
				}
			},
		})
}
