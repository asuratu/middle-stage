package logic

import (
	"context"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.TokenResp, error) {
	getUserByMobileLogic := NewGetUserByMobileLogic(l.ctx, l.svcCtx)
	_, err := getUserByMobileLogic.GetUserByMobile(&user.GetUserByMobileReq{
		Phone: in.Phone,
	})
	// 用户存在
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PhoneRegistered), "phone:%s", in.Phone)
	}

	return &user.TokenResp{}, nil
}
