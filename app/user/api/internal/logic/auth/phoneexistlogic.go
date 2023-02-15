package auth

import (
	"context"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/pkg/errors"

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
	userRsp := &user.GetUserResp{}
	userRsp, err = l.svcCtx.UserRpc.GetUserByMobile(l.ctx, &user.GetUserByMobileReq{
		Phone: req.Phone,
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotFound), "Phone : %+v , err: %v", req.Phone, err)
	}

	return &types.PhoneExistReply{
		Exist: userRsp.User.Id != 0,
	}, nil
}
