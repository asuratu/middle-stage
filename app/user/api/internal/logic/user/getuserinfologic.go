package user

import (
	"context"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/common/ctxdata"

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
	logx.Infof("GetUserInfo")
	// 1. 获取用户id
	uid := ctxdata.GetUidFromCtx(l.ctx)
	logx.Infof("uid: %d", uid)
	// 2. 获取用户信息
	userName := ctxdata.GetUserNameFromCtx(l.ctx)
	logx.Infof("userName: %s", userName)

	// 3. 返回用户信息

	return
}
