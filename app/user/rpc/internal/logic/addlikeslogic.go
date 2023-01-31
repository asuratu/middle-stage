package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLikesLogic {
	return &AddLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------likes-----------------------
func (l *AddLikesLogic) AddLikes(in *user.AddLikesReq) (*user.AddLikesResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddLikesResp{}, nil
}
