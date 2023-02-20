package auth

import (
	"net/http"

	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewLoginByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginByPhone(&req)
		result.HttpResult(r, w, resp, err)
	}
}
