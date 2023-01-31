package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLinksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLinksLogic {
	return &SearchLinksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLinksLogic) SearchLinks(in *user.SearchLinksReq) (*user.SearchLinksResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchLinksResp{}, nil
}
