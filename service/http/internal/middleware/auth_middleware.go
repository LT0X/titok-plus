package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/http/internal/config"
)

type AuthMiddleware struct {
	AccessSecret string
	AccessExpire int64
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		AccessSecret: c.Auth.AccessSecret,
		AccessExpire: c.Auth.AccessExpire,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var token string
		if token = r.URL.Query().Get("token"); token == "" {
			token = r.PostForm.Get("token")
			if token == "" {
				token = r.FormValue("token")
			}
		}

		//账号未登录
		if token == "" {
			httpx.OkJson(w, apiErr.NotLogin)
			return
		}

		//开始鉴权
		ok, err := utils.ValidToken(token, m.AccessSecret)
		if ok || err != nil {
			httpx.OkJson(w, apiErr.InvalidToken)
			return
		}

		next(w, r)
	}
}
