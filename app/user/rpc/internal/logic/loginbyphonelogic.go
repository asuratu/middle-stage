package logic

import (
	"context"
	"middle/common/xerr"
	"middle/pkg/hash"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(in *user.LoginByPhoneReq) (*user.TokenResp, error) {
	getUserByMobileLogic := NewGetUserByMobileLogic(l.ctx, l.svcCtx)
	userModel, err := getUserByMobileLogic.GetUserByMobile(&user.GetUserByMobileReq{
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	// 用户不存在
	if userModel == nil || userModel.User.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PhoneNotRegistered), "phone:%s", in.Phone)
	}

	// 检查密码
	if !hash.BcryptCheck(in.Phone, userModel.User.Password) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PasswordError), "phone:%s", in.Phone)
	}

	// 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userModel.User.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userModel.User.Id)
	}

	return &user.TokenResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
