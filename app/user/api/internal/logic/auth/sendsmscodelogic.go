package auth

import (
	"context"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/pkg/errors"

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
	rsp := &user.VerifyCaptchaResp{}
	rsp, err = l.svcCtx.UserRpc.VerifyCaptcha(l.ctx, &user.VerifyCaptchaReq{
		CaptchaId:    req.CaptchaId,
		CaptchaImage: req.CaptchaAnswer,
	})
	if err != nil {
		return nil, err
	}

	if rsp.IsOk == false {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GraphCaptchaError), "CaptchaId : %+v , err: %v", req.CaptchaId, err)
	}

	// 2. 发送短信验证码
	_, err = l.svcCtx.UserRpc.SendSmsCode(l.ctx, &user.SendSmsCodeReq{
		Phone: req.Phone,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.SendSmsCodeReply{
		Success: true,
	}

	return
}
