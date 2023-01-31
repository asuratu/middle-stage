package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCategoriesLogic {
	return &DelCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCategoriesLogic) DelCategories(in *user.DelCategoriesReq) (*user.DelCategoriesResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelCategoriesResp{}, nil
}
