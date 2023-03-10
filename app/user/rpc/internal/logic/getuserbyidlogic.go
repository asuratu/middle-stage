package logic

import (
	"context"
	"encoding/json"
	"middle/app/mqueue/cmd/job/jobtype"
	"middle/app/user/rpc/model"
	"middle/common/xerr"

	"github.com/hibiken/asynq"

	"github.com/jinzhu/copier"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdReq) (*user.GetUserResp, error) {
	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user record fail"), "get user record fail FindOne err : %v , Id:%d", err, in.Id)
	}
	resp := &user.User{}
	if userModel == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotFound), "The user does not exist : %v , Id:%d", err, in.Id)
	}
	_ = copier.Copy(&resp, userModel)

	// asynq 队列
	payload, err := json.Marshal(jobtype.TestQueuePayload{
		Name: userModel.Name,
		Time: userModel.CreatedAt,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.JsonMarshalError), "json marshal fail : %v ", err)
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.JobTestQueue, payload), asynq.Queue("critical"))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.AsynqEnqueueError), "asynq enqueue fail : %v ", err)
	}
	return &user.GetUserResp{
		User: resp,
	}, nil
}
