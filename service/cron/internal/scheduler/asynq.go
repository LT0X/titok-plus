package scheduler

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"tiktok-plus/common/cron"
	"tiktok-plus/service/cron/internal/svc"
)

type AsynqServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAsynqServer(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqServer {
	return &AsynqServer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AsynqServer) Start() {
	fmt.Println("AsynqTask start")

	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     l.svcCtx.Config.Redis.Address,
			Password: l.svcCtx.Config.Redis.Password},
		nil,
	)

	syncVideoInfoCacheTask, _ := cron.NewSyncVideoInfoTask()
	entryID, err := scheduler.Register("@every 1h", syncVideoInfoCacheTask)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)
	//
	//syncVideoInfoCacheTask := cron.NewSyncVideoInfoCacheTask()
	//entryID, err = scheduler.Register("@every 301s", syncVideoInfoCacheTask)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}

func (l *AsynqServer) Stop() {
	fmt.Println("AsynqTask stop")
}
