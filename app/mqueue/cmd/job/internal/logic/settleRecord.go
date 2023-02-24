package logic

import (
	"context"
	"middle/app/mqueue/cmd/job/internal/svc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/hibiken/asynq"
)

type SettleRecordHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSettleRecordHandler(svcCtx *svc.ServiceContext) *SettleRecordHandler {
	return &SettleRecordHandler{
		svcCtx: svcCtx,
	}
}

func (l *SettleRecordHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// 打印当前时间
	now := time.Now().Format("2006-01-02 15:04:05")
	logx.Infof("schedule job demo -----> every 3 second exec , now: %s", now)
	return nil
}
