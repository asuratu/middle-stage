// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"middle/common/globalkey"
	"middle/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheMsUserUserIdPrefix = "cache:msUser:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneWithTrashed(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) error
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id           int64          `db:"id"`
		Name         string         `db:"name"`
		Email        sql.NullString `db:"email"`
		Phone        sql.NullString `db:"phone"`
		Password     sql.NullString `db:"password"`
		CreatedAt    time.Time      `db:"created_at"`
		UpdatedAt    time.Time      `db:"updated_at"`
		City         sql.NullString `db:"city"`
		Introduction sql.NullString `db:"introduction"`
		Avatar       sql.NullString `db:"avatar"`
		DeleteTime   time.Time      `db:"delete_time"`
		DelState     int64          `db:"del_state"`
		Version      int64          `db:"version"` // 版本号
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		} else {
			return conn.ExecCtx(ctx, query, id)
		}
	}, msUserUserIdKey)
	return err
}

// FindOne 只查询未被软删除的数据
func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, msUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneWithTrashed 查询所有数据, 包括被软删除的数据
func (m *defaultUserModel) FindOneWithTrashed(ctx context.Context, id int64) (*User, error) {
	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, msUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {

	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version)
		} else {
			return conn.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version)
		}
	}, msUserUserIdKey)

}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) error {

	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.Id)
		} else {
			return conn.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.Id)
		}
	}, msUserUserIdKey)

	return err
}

func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error {
	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	msUserUserIdKey := fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.Id, oldVersion)
		} else {
			return conn.ExecCtx(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.City, data.Introduction, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.Id, oldVersion)
		}
	},
		msUserUserIdKey)

	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DbUpdateAffectedZeroError)
	}

	return nil
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMsUserUserIdPrefix, primary)
}
func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
