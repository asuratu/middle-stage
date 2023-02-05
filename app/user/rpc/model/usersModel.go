package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		customUsersLogicModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}

	customUsersLogicModel interface {
		FindOneWithExpire(ctx context.Context, id int64, expire time.Duration) (*Users, error)
		FindOneByField(ctx context.Context, filed, value string) (*Users, error)
	}
)

var (
	cacheGormzeroUsersIdExpirePrefix = "cache:gormzero:users:id:expire:"
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *defaultUsersModel) customCacheKeys(data *Users) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *customUsersModel) FindOneWithExpire(ctx context.Context, id int64, expire time.Duration) (*Users, error) {
	gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdExpirePrefix, id)
	var resp Users
	err := m.QueryWithExpireCtx(ctx, &resp, gormzeroUsersIdKey, expire, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneByField 根据 filed 查询
func (m *defaultUsersModel) FindOneByField(ctx context.Context, filed, value string) (*Users, error) {
	msUserUsersIdKey := fmt.Sprintf("%s%v", cacheMsUserUsersFiledPrefix, filed+value)
	var resp Users
	err := m.QueryCtx(ctx, &resp, msUserUsersIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where(fmt.Sprintf("`%s` = ?", filed), value).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
