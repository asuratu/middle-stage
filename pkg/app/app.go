// Package app 应用信息
package app

import (
	"time"
)

// TimeNowInTimezone 获取当前时间，支持时区
func TimeNowInTimezone() time.Time {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(timezone)
}

// URL 传参 path 拼接站点的 URL
func URL(url, path string) string {
	return url + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(url, path string) string {
	return URL(url, "/v1/"+path)
}
