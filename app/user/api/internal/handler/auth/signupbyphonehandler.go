package auth

import (
	"net/http"

	"middle/app/user/api/internal/logic/auth"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/common/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignupByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignupByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		ctx := gctx.New()
		if err := g.Validator().Data(req).Run(ctx); err != nil {
			result.ValidateErrorResult(r, w, err)
			return
		}

		l := auth.NewSignupByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.SignupByPhone(&req)
		result.HttpResult(r, w, resp, err)
	}
}
