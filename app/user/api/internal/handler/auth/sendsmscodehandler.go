package auth

import (
	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"middle/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendSmsCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		ctx := gctx.New()
		if err := g.Validator().Data(req).Run(ctx); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewSendSmsCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendSmsCode(&req)
		result.HttpResult(r, w, resp, err)
	}
}
