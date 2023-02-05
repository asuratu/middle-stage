package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByIdLogic {
	return &GetUsersByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersByIdLogic) GetUsersById(in *user.GetUsersByIdReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserResp{}, nil
}
