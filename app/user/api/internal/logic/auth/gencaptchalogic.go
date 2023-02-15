package auth

import (
	"context"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/app/user/rpc/user"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenCaptchaLogic {
	return &GenCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenCaptchaLogic) GenCaptcha() (resp *types.CaptchaReply, err error) {
	rsp := &user.SendImageCaptchaResp{}
	rsp, err = l.svcCtx.UserRpc.SendImageCaptcha(l.ctx, &user.SendImageCaptchaReq{})
	if err != nil {
		return nil, err
	}
	resp = &types.CaptchaReply{}
	err = copier.Copy(resp, rsp)
	if err != nil {
		return nil, err
	}
	return
}
