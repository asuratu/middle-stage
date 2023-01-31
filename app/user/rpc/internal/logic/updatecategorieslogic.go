package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoriesLogic {
	return &UpdateCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoriesLogic) UpdateCategories(in *user.UpdateCategoriesReq) (*user.UpdateCategoriesResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateCategoriesResp{}, nil
}
