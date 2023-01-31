package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLinksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLinksLogic {
	return &UpdateLinksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLinksLogic) UpdateLinks(in *user.UpdateLinksReq) (*user.UpdateLinksResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateLinksResp{}, nil
}
