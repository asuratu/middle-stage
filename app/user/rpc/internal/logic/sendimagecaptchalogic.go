package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendImageCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendImageCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendImageCaptchaLogic {
	return &SendImageCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendImageCaptchaLogic) SendImageCaptcha(in *user.SendImageCaptchaReq) (*user.SendImageCaptchaResp, error) {
	// todo: add your logic here and delete this line

	return &user.SendImageCaptchaResp{}, nil
}
