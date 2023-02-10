package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCategoryLogic {
	return &SearchCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCategoryLogic) SearchCategory(in *user.SearchCategoryReq) (*user.SearchCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchCategoryResp{}, nil
}
