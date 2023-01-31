package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTopicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTopicsLogic {
	return &SearchTopicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTopicsLogic) SearchTopics(in *user.SearchTopicsReq) (*user.SearchTopicsResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchTopicsResp{}, nil
}
