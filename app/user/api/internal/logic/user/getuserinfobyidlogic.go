package user

import (
	"context"
	"middle/app/user/rpc/user"

	"github.com/jinzhu/copier"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(req *types.UserInfoByIdReq) (resp *types.UserInfoReply, err error) {
	userRsp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoReply{}
	err = copier.Copy(resp, userRsp.User)
	return
}
