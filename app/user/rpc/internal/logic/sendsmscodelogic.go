package logic

import (
	"context"
	"middle/app/user/rpc/pkg/verifycode"
	"middle/common/xerr"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsCodeLogic {
	return &SendSmsCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsCodeLogic) SendSmsCode(in *user.SendSmsCodeReq) (*user.SendSmsCodeResp, error) {
	if ok := verifycode.NewVerifyCode(l.svcCtx.Config, l.svcCtx.Redis).SendSMS(in.Phone); !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SendSmsError), "phone: %s", in.Phone)
	}
	return &user.SendSmsCodeResp{}, nil
}
