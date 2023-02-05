package logic

import (
	"context"
	"github.com/pkg/errors"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/model"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrInvalidParams = xerr.NewErrMsg("Invalid params")

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

	logx.Infof("GetUsersByAccountLogic.GetUsersByAccount in:%+v", in)

	// phone and email is empty
	if in.Phone == "" && in.Email == "" {
		return nil, errors.Wrap(ErrInvalidParams, "phone and email is empty")
	}

	// phone 不为空, 通过 phone 查询
	userModel := new(model.Users)
	err := error(nil)
	if in.Phone != "" {
		userModel, err = l.svcCtx.UserModel.FindOneByField(l.ctx, "phone", in.Phone)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the record"), "Failed to query the record  rpc GetUsersByAccount fail , phone : %s , err : %v", in.Phone, err)
		}
	}

	// email 不为空, 通过 email 查询
	if in.Email != "" {
		userModel, err = l.svcCtx.UserModel.FindOneByField(l.ctx, "email", in.Email)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the record"), "Failed to query the record  rpc GetUsersByAccount fail , email : %s , err : %v", in.Email, err)
		}
	}

	// 返回结果
	return &user.GetUserResp{
		Users: &user.Users{
			Id:   userModel.Id,
			Name: userModel.Name,
		},
	}, nil
}
