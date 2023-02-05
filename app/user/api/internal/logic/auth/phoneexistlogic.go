package auth

import (
	"context"
	"middle/app/user/rpc/userclient"

	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneExistLogic {
	return &PhoneExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneExistLogic) PhoneExist(req *types.PhoneExistReq) (resp *types.PhoneExistReply, err error) {
	logx.Infof("phone exist: %s", req.Phone)

	_, err = l.svcCtx.UserRpc.GetUsersByAccount(l.ctx, &userclient.GetUsersByAccountReq{
		Phone: req.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &types.PhoneExistReply{
		Exist: true,
	}, nil
}
