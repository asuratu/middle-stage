package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLikesLogic {
	return &DelLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLikesLogic) DelLikes(in *user.DelLikesReq) (*user.DelLikesResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelLikesResp{}, nil
}
