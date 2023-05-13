package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"tiktok-plus/service/mq/internal/config"
	"tiktok-plus/service/mq/internal/handler"
	"tiktok-plus/service/mq/internal/svc"
)

var configFile = flag.String("f", "etc/mq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	logx.DisableStat()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	serviceGroup.Add(handler.NewAsynqServer(ctx, svcContext))
	serviceGroup.Start()

}
