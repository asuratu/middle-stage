package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUsersLogic {
	return &AddUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------users-----------------------

func (l *AddUsersLogic) AddUsers(in *user.AddUsersReq) (*user.AddUsersResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddUsersResp{}, nil
}
