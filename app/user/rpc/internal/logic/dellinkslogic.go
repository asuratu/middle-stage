package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLinksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLinksLogic {
	return &DelLinksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLinksLogic) DelLinks(in *user.DelLinksReq) (*user.DelLinksResp, error) {
	// todo: add your logic here and delete this line

	return &user.DelLinksResp{}, nil
}
