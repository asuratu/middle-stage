// Package helpers 存放辅助方法
package helpers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	mathrand "math/rand"
	"reflect"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString 生成长度为 length 的随机字符串
// 此函数不是加密安全的，不应用于生成密钥或其他敏感数据
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}

// SecureRandomString 生成长度为 length 的随机字符串
// 此函数是加密安全的，可以用于生成密钥或其他敏感数据
// 在使用此函数生成密钥或其他敏感数据时，应根据具体的应用场景和要求进行必要的安全验证和保护
func SecureRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

// StringToNullString 将 string 转为 sql.NullString
func StringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  !Empty(s),
	}
}

// RandomMobile 生成随机手机号码
func RandomMobile() string {
	return fmt.Sprintf("1%s", RandomNumber(10))
}

// CostTime 计算耗时, 单位为毫秒
func CostTime(start time.Time) int64 {
	return time.Since(start).Milliseconds()
}

// CostTimeMin 计算耗时, 单位为分钟
func CostTimeMin(start time.Time) float64 {
	return time.Since(start).Minutes()
}
