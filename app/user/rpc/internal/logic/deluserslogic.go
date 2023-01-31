package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUsersLogic {
	return &DelUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUsersLogic) DelUsers(in *user.DelUsersReq) (*user.DelUsersResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelUsersResp{}, nil
}
