package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCategoryLogic {
	return &DelCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCategoryLogic) DelCategory(in *user.DelCategoryReq) (*user.DelCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelCategoryResp{}, nil
}
