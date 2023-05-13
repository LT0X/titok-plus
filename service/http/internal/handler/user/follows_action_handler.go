package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok-plus/service/http/internal/logic/user"
	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"
)

func FollowsActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowsActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFollowsActionLogic(r.Context(), svcCtx)
		resp, err := l.FollowsAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
