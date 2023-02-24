package logic

import (
	"context"
	"encoding/json"

	"middle/app/mqueue/cmd/job/internal/svc"
	"middle/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// TestQueueHandler test queue job
type TestQueueHandler struct {
	svcCtx *svc.ServiceContext
}

func NewTestQueueHandler(svcCtx *svc.ServiceContext) *TestQueueHandler {
	return &TestQueueHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask every one minute exec : if return err != nil , asynq will retry
func (l *TestQueueHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.TestQueuePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(err, "unmarshal payload")
	}
	logx.Infof("-----------------> handle test queue job , payload: %+v", p)
	return nil
}
