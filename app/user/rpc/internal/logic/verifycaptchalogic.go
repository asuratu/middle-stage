package logic

import (
	"context"
	"middle/app/user/rpc/pkg/captcha"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(in *user.VerifyCaptchaReq) (*user.VerifyCaptchaResp, error) {
	result := true
	// 图片验证码
	if ok := captcha.NewCaptcha(l.svcCtx.Config, l.svcCtx.Redis).VerifyCaptcha(in.CaptchaId, in.CaptchaImage); !ok {
		result = false
	}
	return &user.VerifyCaptchaResp{
		IsOk: result,
	}, nil
}
