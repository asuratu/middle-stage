package user

import (
	"net/http"

	"middle/app/user/api/internal/logic/user"
	"middle/app/user/api/internal/svc"
	"middle/common/result"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		result.HttpResult(r, w, resp, err)
	}
}
