package logic

import (
	"fmt"
	"middle/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// scheduler job ------> go-zero-middle/app/mqueue/cmd/job/internal/logic/settleRecord.go.
func (l *MqueueScheduler) settleRecordScheduler() {

	task := asynq.NewTask(jobtype.ScheduleSettleRecord, nil)
	// every one minute exec
	//entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	// every 3 seconds exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 3s", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【settleRecordScheduler】 registered an  entry: %q \n", entryID)
}
