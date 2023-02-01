package users

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"middle/app/user/api/internal/logic/users"
	"middle/app/user/api/internal/svc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := users.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
