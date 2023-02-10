package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTopicLogic {
	return &DelTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTopicLogic) DelTopic(in *user.DelTopicReq) (*user.DelTopicResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelTopicResp{}, nil
}
