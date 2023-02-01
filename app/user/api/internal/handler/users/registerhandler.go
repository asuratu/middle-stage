package users

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"middle/app/user/api/internal/logic/users"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := users.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
