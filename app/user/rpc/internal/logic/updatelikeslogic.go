package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLikesLogic {
	return &UpdateLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLikesLogic) UpdateLikes(in *user.UpdateLikesReq) (*user.UpdateLikesResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateLikesResp{}, nil
}
