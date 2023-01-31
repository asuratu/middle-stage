package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTopicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTopicsLogic {
	return &DelTopicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTopicsLogic) DelTopics(in *user.DelTopicsReq) (*user.DelTopicsResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelTopicsResp{}, nil
}
