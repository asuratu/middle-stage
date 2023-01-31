package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoriesLogic {
	return &AddCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------categories-----------------------
func (l *AddCategoriesLogic) AddCategories(in *user.AddCategoriesReq) (*user.AddCategoriesResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddCategoriesResp{}, nil
}
