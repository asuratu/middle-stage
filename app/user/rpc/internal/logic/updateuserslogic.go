package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsersLogic {
	return &UpdateUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUsersLogic) UpdateUsers(in *user.UpdateUsersReq) (*user.UpdateUsersResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateUsersResp{}, nil
}
