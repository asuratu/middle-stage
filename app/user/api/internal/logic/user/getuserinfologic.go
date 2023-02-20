package user

import (
	"context"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/app/user/rpc/user"
	"middle/common/ctxdata"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoReply, err error) {
	// 1. 获取用户id
	uid := ctxdata.GetUidFromCtx(l.ctx)
	logx.Infof("uid: %d", uid)

	// 2. 返回用户信息
	userRsp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	// 3. 返回用户信息
	resp = &types.UserInfoReply{}
	err = copier.Copy(resp, userRsp.User)
	if err != nil {
		return nil, err
	}

	return
}
