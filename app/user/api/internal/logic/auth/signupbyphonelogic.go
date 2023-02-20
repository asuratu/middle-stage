package auth

import (
	"context"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/jinzhu/copier"

	"github.com/pkg/errors"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignupByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupByPhoneLogic {
	return &SignupByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupByPhoneLogic) SignupByPhone(req *types.SignupByPhoneReq) (resp *types.TokenReply, err error) {
	// 1. 校验验证码
	rsp := &user.VerifySmsCodeResp{}
	rsp, err = l.svcCtx.UserRpc.VerifySmsCode(l.ctx, &user.VerifySmsCodeReq{
		Phone:   req.Phone,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		return nil, err
	}

	if rsp.IsOk == false {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "Phone : %+v , err: %v", req.Phone, err)
	}

	// 2. 注册
	registerReq := &user.RegisterReq{}
	err = copier.Copy(registerReq, req)
	if err != nil {
		return nil, err
	}
	registerRsp := &user.TokenResp{}
	registerRsp, err = l.svcCtx.UserRpc.Register(l.ctx, registerReq)
	if err != nil {
		return nil, err
	}

	// 3. 返回token
	resp = &types.TokenReply{}
	err = copier.Copy(resp, registerRsp)
	if err != nil {
		return nil, err
	}
	return
}
