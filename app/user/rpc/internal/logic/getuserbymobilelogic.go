package logic

import (
	"context"

	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/model"
	"middle/app/user/rpc/user"
	"middle/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *user.GetUserByMobileReq) (*user.GetUserResp, error) {
	whereBuilder := l.svcCtx.UserModel.RowBuilder().Where(
		"phone = ?",
		in.Phone,
	)

	userModel, err := l.svcCtx.UserModel.FindOneByQuery(l.ctx, whereBuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user record fail"), "get user record fail FindOneByQuery  err : %v , phone:%s", err, in.Phone)
	}

	var resp user.GetUserResp
	if userModel != nil {
		_ = copier.Copy(&resp.User, userModel)
	}

	return &resp, nil
}
