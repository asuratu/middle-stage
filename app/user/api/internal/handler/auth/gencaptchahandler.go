package auth

import (
	"net/http"

	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/common/result"
)

func GenCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewGenCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GenCaptcha()
		result.HttpResult(r, w, resp, err)
	}
}
