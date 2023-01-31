package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinksByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLinksByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinksByIdLogic {
	return &GetLinksByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLinksByIdLogic) GetLinksById(in *user.GetLinksByIdReq) (*user.GetLinksByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetLinksByIdResp{}, nil
}
