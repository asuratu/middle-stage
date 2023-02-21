package logic

import (
	"context"
	"database/sql"
	"middle/app/user/rpc/model"
	"middle/common/xerr"

	"github.com/pkg/errors"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	// 1. 查找用户信息
	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotFound), "UserNotFound to update user err : %v , in :%+v", "uid != in.Id", in)
		}
		return nil, err
	}

	// 2. 更新用户信息
	userModel.Name = in.Name
	userModel.City = sql.NullString{String: in.City, Valid: true}
	userModel.Introduction = sql.NullString{String: in.Introduction, Valid: true}
	userModel.Avatar = sql.NullString{String: in.Avatar, Valid: true}

	if err := l.svcCtx.UserModel.Update(l.ctx, nil, userModel); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed to update user err : %v , in :%+v", err, in)
	}
	return &user.UpdateUserResp{}, nil

}
