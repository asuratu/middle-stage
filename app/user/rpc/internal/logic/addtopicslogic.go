package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicsLogic {
	return &AddTopicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------topics-----------------------
func (l *AddTopicsLogic) AddTopics(in *user.AddTopicsReq) (*user.AddTopicsResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddTopicsResp{}, nil
}
