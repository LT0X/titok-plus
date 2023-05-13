package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok-plus/service/http/internal/config"
	"tiktok-plus/service/http/internal/middleware"
	"tiktok-plus/service/rpc/user/userclient"
	"tiktok-plus/service/rpc/video/videoclient"
)

type ServiceContext struct {
	Config   config.Config
	Auth     rest.Middleware
	UserRpc  userclient.User
	VideoRpc videoclient.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Auth:     middleware.NewAuthMiddleware(c).Handle,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		VideoRpc: videoclient.NewVideo(zrpc.MustNewClient(c.VideoRPC)),
	}
}
