package video

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok-plus/service/http/internal/logic/video"
	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"
)

func PublishVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewPublishVideoListLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
