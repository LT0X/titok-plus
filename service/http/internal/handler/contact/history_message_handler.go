package contact

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok-plus/service/http/internal/logic/contact"
	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"
)

func HistoryMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HistoryMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contact.NewHistoryMessageLogic(r.Context(), svcCtx)
		resp, err := l.HistoryMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}

}
