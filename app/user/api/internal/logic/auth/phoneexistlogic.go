package auth

import (
	"context"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneExistLogic {
	return &PhoneExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneExistLogic) PhoneExist(req *types.PhoneExistReq) (resp *types.PhoneExistReply, err error) {

	return
}
