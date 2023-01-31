package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikesByIdLogic {
	return &GetLikesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikesByIdLogic) GetLikesById(in *user.GetLikesByIdReq) (*user.GetLikesByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetLikesByIdResp{}, nil
}
