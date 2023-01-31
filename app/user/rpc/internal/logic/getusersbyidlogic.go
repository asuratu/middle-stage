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

func (l *GetUsersByIdLogic) GetUsersById(in *user.GetUsersByIdReq) (*user.GetUsersByIdResp, error) {
	// 打印测试日志
	l.Logger.Infof("GetUsersById: %v", in)
	logx.Infof("************GetUsersById: %v", in)

	// 获取 user 信息
	//if user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id); err != nil {
	//	return nil, err
	//} else {
	//	return &user.GetUsersByIdResp{
	//		Id:       user.Id,
	//		Username: user.Username,
	//		Password: user.Password,
	//	}, nil
	//}

	return &user.GetUsersByIdResp{}, nil
}
