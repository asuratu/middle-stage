package auth

import (
	"github.com/go-playground/validator/v10"
	"middle/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
)

func PhoneExistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneExistReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewPhoneExistLogic(r.Context(), svcCtx)
		resp, err := l.PhoneExist(&req)
		result.HttpResult(r, w, resp, err)
	}
}
