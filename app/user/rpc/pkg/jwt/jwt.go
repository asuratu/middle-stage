// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"middle/app/user/rpc/internal/config"
	"middle/pkg/app"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/service"

	jwtpkg "github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 AccessSecret
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration

	Config config.Config
}

// CustomClaims 自定义载荷
type CustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.RegisteredClaims
}

type TokenRsp struct {
	Token               string `json:"token"`
	ExpireAtTime        int64  `json:"expire_at_time"`
	RefreshExpireAtTime int64  `json:"refresh_expire_at_time"`
}

func NewJWT(c config.Config) *JWT {
	return &JWT{
		SignKey:    []byte(c.JwtAuth.AccessSecret),
		MaxRefresh: time.Duration(c.JwtAuth.RefreshExpire) * time.Minute,
		Config:     c,
	}
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(tokenStr string) (*TokenRsp, error) {
	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenStr)

	// 2. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return nil, err
		}
	}

	// 3. 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)

	// 4. 检查是否过了『最大允许刷新的时间』
	x := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Time.Unix() > x {
		// 4.1 修改过期时间
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())

		// 4.2 根据 claims 生成token对象
		token, err := jwt.createToken(*claims)
		if err != nil {
			logx.Error(err)
			return nil, err
		}

		return &TokenRsp{
			Token:               token,
			ExpireAtTime:        claims.RegisteredClaims.ExpiresAt.Unix(),
			RefreshExpireAtTime: claims.RegisteredClaims.ExpiresAt.Add(jwt.MaxRefresh).Unix(),
		}, nil
	}

	return nil, ErrTokenExpiredMaxRefresh
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(userID int64, userName string) (*TokenRsp, error) {
	c := jwt.Config
	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.expireAtTime()
	claims := CustomClaims{
		strconv.FormatInt(userID, 10),
		userName,
		expireAtTime.Unix(),
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 签名生效时间
			IssuedAt:  jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(expireAtTime),            // 签名过期时间
			Issuer:    c.Name,                                         // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &TokenRsp{
		Token:               token,
		ExpireAtTime:        expireAtTime.Unix(),
		RefreshExpireAtTime: expireAtTime.Add(jwt.MaxRefresh).Unix(),
	}, nil
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims CustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() time.Time {
	timeNow := app.TimeNowInTimezone()

	c := jwt.Config
	var expireTime int64
	// 如果是开发环境，使用 debug 的过期时间
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		expireTime = c.JwtAuth.DebugExpire
	} else {
		expireTime = c.JwtAuth.AccessExpire
	}

	expire := time.Duration(expireTime) * time.Second
	return timeNow.Add(expire)
}
