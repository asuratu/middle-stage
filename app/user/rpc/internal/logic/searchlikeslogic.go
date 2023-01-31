package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLikesLogic {
	return &SearchLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLikesLogic) SearchLikes(in *user.SearchLikesReq) (*user.SearchLikesResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchLikesResp{}, nil
}
