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

type VerifySmsCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifySmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifySmsCodeLogic {
	return &VerifySmsCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifySmsCodeLogic) VerifySmsCode(in *user.VerifySmsCodeReq) (*user.VerifySmsCodeResp, error) {
	// key 和 answer 不能为空
	if in.Phone == "" || in.SmsCode == "" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ParamValidateError), "phone: %s", in.Phone)
	}

	if ok := verifycode.NewVerifyCode(l.svcCtx.Config, l.svcCtx.Redis).CheckAnswer(in.Phone, in.SmsCode); !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "phone: %s", in.Phone)
	}
	return &user.VerifySmsCodeResp{
		IsOk: true,
	}, nil
}
