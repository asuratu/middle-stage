package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByIdLogic {
	return &GetTopicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopicByIdLogic) GetTopicById(in *user.GetTopicByIdReq) (*user.GetTopicByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetTopicByIdResp{}, nil
}
