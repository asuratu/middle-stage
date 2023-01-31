package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCategoriesLogic {
	return &SearchCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCategoriesLogic) SearchCategories(in *user.SearchCategoriesReq) (*user.SearchCategoriesResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchCategoriesResp{}, nil
}
