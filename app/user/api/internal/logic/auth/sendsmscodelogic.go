package auth

import (
	"context"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsCodeLogic {
	return &SendSmsCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsCodeLogic) SendSmsCode(req *types.SendSmsCodeReq) (resp *types.SendSmsCodeReply, err error) {
	// 1. 校验图形验证码

	// 2. 发送短信验证码

	return
}
