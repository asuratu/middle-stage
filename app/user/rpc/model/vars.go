package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// del_state 0:正常 1:已删除
const (
	DelStateNormal = 0
	DelStateDelete = 1
)
