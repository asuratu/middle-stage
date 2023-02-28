package logic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"

	"middle/app/mqueue/cmd/job/internal/svc"
	"middle/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

// RegisterHandler register userinfo to es
type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask every one minute exec : if return err != nil , asynq will retry
func (l *RegisterHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.RegisterUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(err, "unmarshal payload")
	}
	logx.Infof("-----------------> handle register job , payload: %+v", p)
	return nil
}
