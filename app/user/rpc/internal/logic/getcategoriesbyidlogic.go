package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoriesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesByIdLogic {
	return &GetCategoriesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoriesByIdLogic) GetCategoriesById(in *user.GetCategoriesByIdReq) (*user.GetCategoriesByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetCategoriesByIdResp{}, nil
}
