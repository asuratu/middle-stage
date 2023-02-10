package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTopicLogic {
	return &SearchTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTopicLogic) SearchTopic(in *user.SearchTopicReq) (*user.SearchTopicResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchTopicResp{}, nil
}
