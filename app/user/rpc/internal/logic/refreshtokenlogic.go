package logic

import (
	"context"
	jwtpkg "middle/app/user/rpc/pkg/jwt"
	"middle/common/xerr"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *user.RefreshTokenReq) (*user.TokenResp, error) {
	logx.Info("RefreshTokenLogic RefreshToken in: %v", in)
	j := jwtpkg.NewJWT(l.svcCtx.Config)
	tokenRsp, err := j.RefreshToken(in.Token)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.RefreshTokenError), "err: %d", err)
	}

	return &user.TokenResp{
		AccessToken:  tokenRsp.Token,
		AccessExpire: tokenRsp.ExpireAtTime,
		RefreshAfter: tokenRsp.RefreshExpireAtTime,
	}, nil
}
