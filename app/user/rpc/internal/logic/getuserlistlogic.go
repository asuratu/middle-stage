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

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.GetUserListReq) (*user.GetUserListResp, error) {
	var orderBy string
	if in.Sort != "" && in.Order != "" {
		orderBy = in.Sort + " " + in.Order
	}
	list, err := l.svcCtx.UserModel.FindPageListByPage(l.ctx, l.svcCtx.UserModel.RowBuilder(), in.Page, in.PageSize, orderBy)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed to get user list err : %v , in :%+v", err, in)
	}

	resp := &user.GetUserListResp{}
	if list != nil {
		_ = copier.Copy(&resp.List, list)
	}

	return resp, nil
}
