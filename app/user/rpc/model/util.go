// Package models 模型通用方法
package models

import (
	"github.com/spf13/cast"
)

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
