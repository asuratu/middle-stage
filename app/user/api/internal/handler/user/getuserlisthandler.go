package user

import (
	"net/http"

	"middle/app/user/api/internal/logic/user"
	"middle/app/user/api/internal/svc"
	"middle/app/user/api/internal/types"
	"middle/common/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		ctx := gctx.New()
		if err := g.Validator().Data(req).Run(ctx); err != nil {
			result.ValidateErrorResult(r, w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
