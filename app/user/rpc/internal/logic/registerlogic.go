package logic

import (
	"context"
	"database/sql"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/model"
	"middle/app/user/rpc/user"
	"middle/common/tool"
	"middle/common/xerr"
	"middle/pkg/hash"
	"middle/pkg/helpers"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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
	// 开始时间, 用于计算耗时
	startTime := time.Now()
	getUserByMobileLogic := NewGetUserByMobileLogic(l.ctx, l.svcCtx)
	userModel, err := getUserByMobileLogic.GetUserByMobile(&user.GetUserByMobileReq{
		Phone: in.Phone,
	})
	// 计算耗时
	logx.Infof("------------- GetUserByMobile耗时: %v", helpers.CostTime(startTime))
	if err != nil {
		return nil, err
	}
	// 用户存在
	if userModel != nil && userModel.User.Id > 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PhoneRegistered), "phone:%s", in.Phone)
	}

	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 组装数据
		userModel := new(model.User)
		userModel.Phone = sql.NullString{String: in.Phone, Valid: true}
		if len(in.Name) == 0 {
			in.Name = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		userModel.Name = in.Name
		// 计算耗时
		logx.Infof("------------- 组装数据耗时1: %v", helpers.CostTime(startTime))
		if len(in.Password) > 0 {
			userModel.Password = sql.NullString{String: hash.BcryptHash(in.Password), Valid: true}
		}
		// 计算耗时
		logx.Infof("------------- 组装数据耗时2: %v", helpers.CostTime(startTime))
		userModel.City = sql.NullString{String: in.City, Valid: true}
		userModel.Introduction = sql.NullString{String: in.Introduction, Valid: true}
		// 计算耗时
		logx.Infof("------------- 组装数据耗时3: %v", helpers.CostTime(startTime))
		now := carbon.Now().ToStdTime() // 2020-08-05 13:14:15
		userModel.DeleteTime = now
		userModel.CreatedAt = now
		userModel.UpdatedAt = now
		// 计算耗时
		logx.Infof("------------- 组装数据耗时: %v", helpers.CostTime(startTime))
		insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, userModel)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Register db user Insert err:%v,user:%+v", err, userModel)
		}
		// 计算耗时
		logx.Infof("------------- Insert耗时: %v", helpers.CostTime(startTime))
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, userModel)
		}
		userId = lastId
		return nil
	}); err != nil {
		return nil, err
	}

	// 计算耗时
	logx.Infof("------------- Trans耗时: %v", helpers.CostTime(startTime))

	// 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, _ := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	// 计算耗时
	logx.Infof("------------- GenerateToken耗时: %v", helpers.CostTime(startTime))
	return &user.TokenResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
