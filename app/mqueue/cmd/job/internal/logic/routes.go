package logic

import (
	"context"

	"middle/app/mqueue/cmd/job/internal/svc"
	"middle/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// test queue job
	mux.Handle(jobtype.JobTestQueue, NewTestQueueHandler(l.svcCtx))

	// scheduler job
	// mux.Handle(jobtype.ScheduleSettleRecord, NewSettleRecordHandler(l.svcCtx))

	// defer job
	mux.Handle(jobtype.RegisterUserJob, NewRegisterHandler(l.svcCtx))

	// defer job
	// mux.Handle(jobtype.DeferCloseHomestayOrder, NewCloseHomestayOrderHandler(l.svcCtx))

	// queue job , asynq support queue job
	// wait you fill

	return mux
}
