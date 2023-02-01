package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
)

func EmailExistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailExistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := auth.NewEmailExistLogic(r.Context(), svcCtx)
		resp, err := l.EmailExist(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
