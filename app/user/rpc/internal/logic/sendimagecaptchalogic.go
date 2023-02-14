package logic

import (
	"context"
	"middle/app/user/rpc/pkg/captcha"

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
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha(l.svcCtx.Config, l.svcCtx.Redis).GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	if err != nil {
		logx.Error(err)
	}
	// 返回给用户
	return &user.SendImageCaptchaResp{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, nil
}
