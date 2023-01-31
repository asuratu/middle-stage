package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLinksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLinksLogic {
	return &AddLinksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------links-----------------------
func (l *AddLinksLogic) AddLinks(in *user.AddLinksReq) (*user.AddLinksResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddLinksResp{}, nil
}
