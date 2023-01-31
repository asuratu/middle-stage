package logic

import (
	"context"
	"github.com/pkg/errors"
	"middle/app/user/rpc/model"
	"middle/common/xerr"
	"time"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByIdLogic {
	return &GetUsersByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersByIdLogic) GetUsersById(in *user.GetUsersByIdReq) (*user.GetUsersByIdResp, error) {
	// 获取 user 信息
	if u, err := l.svcCtx.UserModel.FindOneWithExpire(l.ctx, in.Id, time.Minute*3); err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotFound), "User Database Exception user : %+v , err: %v", in.Id, err)
		}
		return nil, err
	} else {
		return &user.GetUsersByIdResp{
			Users: &user.Users{
				Id: u.Id,
			},
		}, nil
	}

}
