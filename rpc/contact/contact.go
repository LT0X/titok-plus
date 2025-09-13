package main

import (
	"douyin/rpc/contact/contact"
	"douyin/rpc/contact/internal/cache"
	"douyin/rpc/contact/internal/config"
	"douyin/rpc/contact/internal/server"
	"douyin/rpc/contact/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/contact.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		contact.RegisterContactServer(grpcServer, server.NewContactServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	cache.InitRedis(ctx)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
