// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "middle/app/user/api/internal/handler/auth"
	user "middle/app/user/api/internal/handler/user"
	"middle/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/auth/signup/phone/exist",
				Handler: auth.PhoneExistHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/auth/verify-codes/captcha",
				Handler: auth.GenCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/verify-codes/sms",
				Handler: auth.SendSmsCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/signup/phone/register",
				Handler: auth.SignupByPhoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/login/phone",
				Handler: auth.LoginByPhoneHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/:id",
				Handler: user.GetUserInfoByIdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/userapi/v1"),
	)
}
