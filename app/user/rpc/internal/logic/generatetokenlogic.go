package logic

import (
	"context"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	jwtpkg "middle/app/user/rpc/pkg/jwt"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *user.GenerateTokenReq) (*user.TokenResp, error) {
	j := jwtpkg.NewJWT(l.svcCtx.Config)
	tokenRsp, err := j.IssueToken(in.UserId, in.UserName)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ErrGenerateTokenError), "UserId:%d", in.UserId)
	}

	return &user.TokenResp{
		AccessToken:  tokenRsp.Token,
		AccessExpire: tokenRsp.ExpireAtTime,
		RefreshAfter: tokenRsp.RefreshExpireAtTime,
	}, nil
}
