package logic

import (
	"context"
	"middle/app/user/rpc/model"
	"middle/common/xerr"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByAccountLogic {
	return &GetUsersByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersByAccountLogic) GetUsersByAccount(in *user.GetUsersByAccountReq) (*user.GetUserResp, error) {
	// 获取 user 信息
	if u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id); err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotFound), "User Database Exception user : %+v , err: %v", in.Id, err)
		}
		return nil, err
	} else {
		return &user.GetUserResp{
			Users: &user.Users{
				Id: u.Id,
			},
		}, nil
	}
}
