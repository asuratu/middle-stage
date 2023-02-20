package user

import (
	"net/http"

	"middle/app/user/api/internal/logic/user"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewGetUserInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoById(&req)
		result.HttpResult(r, w, resp, err)
	}
}
