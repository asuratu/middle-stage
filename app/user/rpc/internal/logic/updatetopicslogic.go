package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicsLogic {
	return &UpdateTopicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTopicsLogic) UpdateTopics(in *user.UpdateTopicsReq) (*user.UpdateTopicsResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateTopicsResp{}, nil
}
