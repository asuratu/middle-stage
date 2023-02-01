package auth

import (
	"context"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailExistLogic {
	return &EmailExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailExistLogic) EmailExist(req *types.EmailExistReq) (resp *types.EmailExistReply, err error) {
	// todo: add your logic here and delete this line

	return
}
