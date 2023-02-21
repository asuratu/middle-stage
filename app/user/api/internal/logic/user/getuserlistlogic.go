package user

import (
	"context"
	"middle/app/user/rpc/user"

	"github.com/jinzhu/copier"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListReply, err error) {
	userListReq := &user.GetUserListReq{}
	_ = copier.Copy(userListReq, req)
	userListRsp, err := l.svcCtx.UserRpc.GetUserList(l.ctx, userListReq)
	if err != nil {
		return nil, err
	}
	resp = &types.UserListReply{}
	for _, u := range userListRsp.List {
		simpleUser := &types.SimpleUserInfoReply{}
		_ = copier.Copy(simpleUser, u)
		resp.Users = append(resp.Users, simpleUser)
	}
	return
}
