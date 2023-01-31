package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicsByIdLogic {
	return &GetTopicsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopicsByIdLogic) GetTopicsById(in *user.GetTopicsByIdReq) (*user.GetTopicsByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetTopicsByIdResp{}, nil
}
