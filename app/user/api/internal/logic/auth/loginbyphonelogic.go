package auth

import (
	"context"

	"github.com/jinzhu/copier"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(req *types.LoginByPhoneReq) (resp *types.TokenReply, err error) {
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

	// 2. 登录
	loginReq := &user.LoginByPhoneReq{}
	err = copier.Copy(loginReq, req)
	if err != nil {
		return nil, err
	}
	loginRsp := &user.TokenResp{}
	loginRsp, err = l.svcCtx.UserRpc.LoginByPhone(l.ctx, loginReq)
	if err != nil {
		return nil, err
	}

	// 3. 返回token
	resp = &types.TokenReply{}
	err = copier.Copy(resp, loginRsp)
	if err != nil {
		return nil, err
	}

	return
}
