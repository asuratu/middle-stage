package user

import (
	"context"
	"middle/common/ctxdata"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserReq) (resp *types.UpdateUserReply, err error) {
	// 1. 获取用户ID
	uid := ctxdata.GetUidFromCtx(l.ctx)
	// 2. 更新用户信息
	updateUserReq := new(user.UpdateUserReq)
	_ = copier.Copy(updateUserReq, req)
	updateUserReq.Id = uid
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, updateUserReq)
	if err != nil {
		return nil, err
	}
	return
}
